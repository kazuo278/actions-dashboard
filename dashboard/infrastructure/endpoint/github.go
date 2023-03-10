package endpoint

import (
	"encoding/json"
	"fmt"
	"github/dashboard/model/custom"
	"log"
	"net/http"
)

const baseUrl string = "https://api.github.com/"
const versionHeader string = "X-GitHub-Api-Version"
const version string = "2022-11-28"
const AcceptHeader string = "Accept"
const accept string = "application/vnd.github+json"
const AuthorizationHeader string = "Authorization"

func GetRunnerRegistrationToken(organizationName string, organizationToken string) (custom.GitHubResponse, int, error) {
	// リクエストさの作成
	url := baseUrl + "orgs/" + organizationName + "/actions/runners/registration-token"
	request, _ := http.NewRequest(http.MethodPost, url, nil)
	request.Header.Set(versionHeader, version)
	request.Header.Set(AcceptHeader, accept)
	request.Header.Set(AuthorizationHeader, "Bearer "+organizationToken)
	client := new(http.Client)

	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error: エンドポイント接続時にエラーが発生しました。%s", err)
		return nil, 0, fmt.Errorf("Error: エンドポイント接続時にエラーが発生しました: %w", err)
	}

	defer response.Body.Close()

	// 正常応答
	if response.StatusCode == 201 {
		result := custom.GitHubRegistrationTokenResponse{}
		err = json.NewDecoder(response.Body).Decode(&result)
		if err != nil {
			log.Printf("JSON Decode Error: %s", err)
			return nil, 0, fmt.Errorf("JSON Decode Error: %w", err)
		}
		return &result, response.StatusCode, nil
		// エラー応答
	} else {
		result := custom.GitHubErrorResponse{}
		err = json.NewDecoder(response.Body).Decode(&result)
		if err != nil {
			log.Printf("JSON Decode Error: %s", err)
			return nil, 0, fmt.Errorf("JSON Decode Error: %w", err)
		}
		return &result, response.StatusCode, nil
	}
}
