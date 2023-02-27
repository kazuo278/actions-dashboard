# README

## 使い方

- ダッシュボード
  - http://${DASHBOARD_APP_HOST}/dashboard

- 実行履歴登録
  - ジョブ実行開始履歴の登録

    ```sh
    $ REPOSITORY_NAME=$(cat $GITHUB_EVENT_PATH | jq -r .repository.full_name)
    $ curl -X POST ${DASHBOARD_APP_HOST}/actions/history -H 'Content-Type: application/json' -d @- <<EOM
    {"repository_id":"$GITHUB_REPOSITORY_ID", "repository_name":"$REPOSITORY_NAME", "run_id":"$GITHUB_RUN_ID"}
    EOM
    ```

  - ジョブ終了履歴の登録

    ```sh
    $ curl -X PUT ${DASHBOARD_APP_HOST}/actions/history -H 'Content-Type: application/json' -d @- <<EOM
    {"repository_id":"$GITHUB_REPOSITORY_ID", "run_id":"$GITHUB_RUN_ID"}
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
