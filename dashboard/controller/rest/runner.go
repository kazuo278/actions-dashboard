package rest

import (
	"log"
	"net/http"

	"github/dashboard/application"
	"github/dashboard/model/custom"

	"github.com/labstack/echo/v4"
)

// 実行履歴を登録する
// POST: /actions/runner/registraion-token
// { organization_name: <string>, organization_key: <string>}
func PostRunnerRegistrationToken(c echo.Context) error {
	// JSONリクエストを取得
	body := custom.GitHubRegistrationTokenRequest{}
	c.Bind(&body)
	// 実行履歴を登録
	result, statusCode, err := application.GetRegistrationToken(body.OrganizationName, body.OrganizationKey)
	if err != nil || statusCode == 0 {
		log.Printf(err.Error())
		return c.JSON(http.StatusInternalServerError, "予期しないエラーが発生しました")
	}
	return c.JSON(statusCode, result)
}
