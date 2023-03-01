# README

## 使い方

- ダッシュボード
  - http://${DASHBOARD_APP_HOST}/dashboard

- 実行履歴登録
  - ジョブ実行開始履歴の登録

    ```sh
    $ WORKFLOW_REF=$(echo $GITHUB_WORKFLOW_REF | sed "s%$GITHUB_REPOSITORY/%%")
    $ curl -X POST ${DASHBOARD_APP_HOST}/actions/history -H 'Content-Type: application/json' -d @- <<EOM
    {"repository_id":"$GITHUB_REPOSITORY_ID", "repository_name":"$GITHUB_REPOSITORY", "run_id":"$GITHUB_RUN_ID", "workflow_ref":"$WORKFLOW_REF", "job":"$GITHUB_JOB"}
    EOM
    ```

  - ジョブ終了履歴の登録

    ```sh
    $ curl -X PUT ${DASHBOARD_APP_HOST}/actions/history -H 'Content-Type: application/json' -d @- <<EOM
    {"repository_id":"$GITHUB_REPOSITORY_ID", "run_id":"$GITHUB_RUN_ID"}
    EOM
    ```

## アプリ実行

- コンテナ実行

  ```
  docker compose up
  ```

- ローカル実行(開発モード)

  ```
  docker compose start postgres
  docker compose start pgweb
  DATABASE_URL=postgres://appuser:password@localhost:5432/github-actions?sslmode=disable go run
  ```

- アプリビルド

  ```sh
  cd dashboard
  # linux用
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo
  # ビルド環境用
  go build
  ```

