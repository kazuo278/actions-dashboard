package application

import (
	"bufio"
	"fmt"
	"github/dashboard/infrastructure/endpoint"
	"github/dashboard/model/custom"
	"log"
	"os"
	"strings"
)

const secretPath string = "/run/secrets/"

// ランナートークンを生成する
func GetRegistrationToken(organizationName string, organizationKey string) (custom.GitHubResponse, int, error) {

	organizationToken := ""

	// OrganizationTokenを取得する
	// シークレット(/run/secrets/*)を探す
	fp, err := os.Open(secretPath + organizationKey)
	if err == nil {
		log.Printf("Organizationトークンをシークレットから取得します")
		defer fp.Close()
		scanner := bufio.NewScanner(fp)
		scanner.Scan()
		organizationToken = scanner.Text()
	} else {
		// 見つからない場合は環境変数を探す
		log.Printf("Organizationトークンを環境変数から取得します")
		organizationToken = os.Getenv(strings.ToUpper(organizationKey))
	}

	if organizationToken == "" {
		return nil, 0, fmt.Errorf("Organizationトークンが見つかりませんでした")
	}

	tokenResponse, statusCode, err := endpoint.GetRunnerRegistrationToken(organizationName, organizationToken)
	return tokenResponse, statusCode, err
}
