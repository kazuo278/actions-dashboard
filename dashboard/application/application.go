package application

import (
	"log"
	"time"

	"github/dashboard/infrastructure/database"
	"github/dashboard/model"
	"github/dashboard/model/custom"
)

// 実行履歴を開始ステータスで登録する
func CreateHistoryWithStarted(repositoryId string, repositoryName string, runId string, workflowRef string, jobName string, runAttempt string) *model.History {
	// リポジトリを取得
	repository := database.GetRepositoryById(repositoryId)
	if (*repository == model.Repository{}) {
		log.Print("リポジトリ名" + repositoryName + "は一覧に存在しないため登録します")
		// リポジトリが存在しない(新規Actions実行の)場合、登録
		repository := new(model.Repository)
		repository.RepositoryId = repositoryId
		repository.RepositoryName = repositoryName
		database.CreateRepository(repository)
	}
	history := new(model.History)
	history.RepositoryId = repositoryId
	history.RunId = runId
	history.WorkflowRef = workflowRef
	history.JobName = jobName
	history.RunAttempt = runAttempt
	history.Status = "STARTED"
	history.StartedAt = NowJST()
	database.CreateHistory(history)

	return history
}

// 実行履歴を終了ステータスで更新する
func UpdateHistoryWithFinished(repositoryId string, runId string, jobName string, runAttempt string) *model.History {
	// 更新対象を取得
	history := database.GetHistoryById(repositoryId, runId, jobName, runAttempt)
	// ステータスを変更
	history.FinishedAt = NowJST()
	history.Status = "FINISHED"
	// 更新
	database.UpdateHistory(history)

	return history
}

// 実行履歴を取得する
func GetHistories(limit int, offset int, repositoryId string, repositoryName string, workflowRef string, jobName string, runAttempt string, status string, startedAt string, finishedAt string) *custom.HistoryResponse {
	histories, count := database.GetHistories(limit, offset, repositoryId, repositoryName, workflowRef, jobName, runAttempt, status, startedAt, finishedAt)
	result := new(custom.HistoryResponse)
	result.Histories = histories
	result.Count = count

	return result
}

// リポジトリごとの実行回数を取得する
func GetHistoryCount(repositoryName string, startedAt string, finishedAt string) *custom.HistoryCounterResponse {
	result := new(custom.HistoryCounterResponse)
	result.Counts = database.GetHistryCount(repositoryName, startedAt, finishedAt)
	return result
}

func NowJST() *time.Time {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	nowJST := time.Now().In(jst)
	return &nowJST
}
