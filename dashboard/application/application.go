package application

import (
	"time"
	"log"

	"github/dashboard/infrastructure/database"
	"github/dashboard/model"
	"github/dashboard/model/custom"
)

// 実行履歴を開始ステータスで登録する
func CreateHistoryWithStarted(repositoryId string, repositoryName string, runId string) *model.History {
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
	// 実行履歴を登録
	history := new(model.History)
	history.RepositoryId = repositoryId
	history.RunId = runId
	history.Status = "STARTED"
	history.StartedAt = time.Now()
	database.CreateHistory(history)

	return history
}

// 実行履歴を終了ステータスで更新する
func UpdateHistoryWithFinished(repositoryId string, runId string) *model.History {
	// 更新対象を取得
	history := database.GetHistoryById(repositoryId, runId)
	// ステータスを変更
	history.FinishedAt = time.Now()
	history.Status = "FINISHED"
	// 更新
	database.UpdateHistory(history)

	return history
}

// 実行履歴を取得する
func GetHistories(limit int, offset int, repositoryId string, repositoryName string, workflowRef string, jobName string, status string, startedAt string, finishedAt string) *custom.HistoryResponse {
	histories, count:= database.GetHistories(limit, offset, repositoryId, repositoryName, workflowRef, jobName, status, startedAt, finishedAt)
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