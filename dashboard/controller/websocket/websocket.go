package websocket

import (
	"log"
	"container/list"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

// 実行履歴更新フラグ
var IsUpdated bool = false

// クライアントのコネクションリスト
var conns list.List

func Socket(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		var listElm *list.Element

		// すでに登録済みのコネクションか判定
		contain := false
		for con := conns.Front(); con != nil; con = con.Next() {
			if con.Value == ws {
				contain = true
				break
			}
		}
		// 未登録の場合、connectionリストに追加
		if !contain {
			listElm = conns.PushBack(ws)
			// conns.PushBack(ws)
			log.Print("新たなクライアントが接続されました。クライアント数: ", conns.Len())
		}

		for {
			// 5sおきに接続確認
			// 接続できない場合はコネクションを切断
			time.Sleep(time.Second * 5)
			err := websocket.Message.Send(ws, "CONNECTION CHECK")
			if err != nil {
				c.Logger().Error(err)
				conns.Remove(listElm)
				ws.Close()
				log.Print("受信できないクライアントを削除しました。クライアント数: ", conns.Len())
				break
			}

			// 更新がない場合
			if !IsUpdated {
				continue
			}

			// 更新がある場合、全てのクライアントへ更新を通知
			log.Print("実行履歴が更新されたためクライアントへ通知します。クライアント数: ", conns.Len())
			for con := conns.Front(); con != nil; con = con.Next() {
				err := websocket.Message.Send(con.Value.(*websocket.Conn), "UPDATED")
				if err != nil {
					c.Logger().Error(err)
					// errorが発生するとループが止まるので、接続ごとにコネクションクローズするように変更
					// con.Value.(*websocket.Conn).Close()
					// conns.Remove(con)
					log.Print("正常に送信できないクライアントが存在しました。")
				}
			}
			IsUpdated = false
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}