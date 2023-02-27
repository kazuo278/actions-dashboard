# README

## 使い方

- ダッシュボード
  - http://${HOST}/dashboard

- 実行履歴登録
  - ジョブ実行開始履歴の登録

    ```sh
    $ curl -X POST ${HOST}/actions/history -H 'Content-Type: application/json' -d @- <<EOM
    {"repository_id":"$REPOSITORY_ID", "repository_name":"$REPOSITORY_NAME", "run_id":"$RUN_ID"}
    EOM
    ```

  - ジョブ終了履歴の登録

    ```sh
    $ curl -X PUT ${HOST}/actions/history -H 'Content-Type: application/json' -d @- <<EOM
    {"repository_id":"$REPOSITORY_ID", "run_id":"$RUN_ID"}
    EOM
    ```

## アプリ実行

- ビルド

  ```sh
  $ cd dashboard
  $ go CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo
  ```

- 実行

  ```sh
  $ cd dashboard
  $ ./dashboard
  ```
