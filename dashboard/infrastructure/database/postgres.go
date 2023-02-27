package database

import (
	"os"
	"log"

	"github/dashboard/model"
	"github/dashboard/model/custom"
	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(){
	connectionUrl := os.Getenv("DATABASE_URL")
	var err error
	sqlDB, err := sql.Open("pgx", connectionUrl)
	if err != nil {
		log.Print("DB接続に失敗しました")
		panic(err)
	}
	db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Print("DB接続に失敗しました")
		panic(err)
	}
}

func GetHistories(limit int, offset int, repositoryId string, repositoryName string, status string, startedAt string, finishedAt string) (*[]custom.HistoryRepository, int) {
	result := []custom.HistoryRepository{}
	count := 0
	sql := db.Table("histories").Select("histories.*, repositories.repository_name").Joins("left join repositories on repositories.repository_id = histories.repository_id")
	countSql := db.Table("histories").Select("count(1)").Joins("left join repositories on repositories.repository_id = histories.repository_id")

	if repositoryId != "" {
		sql.Where("histories.repository_id = ?", repositoryId)
		countSql.Where("histories.repository_id = ?", repositoryId)
	}

	if repositoryName != "" {
		sql.Where("repository_name LIKE ?", repositoryName + "%")
		countSql.Where("repository_name LIKE ?", repositoryName + "%")
	}

	if status != "" {
		sql.Where("status = ?", status)
		countSql.Where("status = ?", status)
	}

	if startedAt != "" {
		sql.Where("started_at >= ?", startedAt)
		countSql.Where("started_at >= ?", startedAt)
	}

	if finishedAt != "" {
		sql.Where("finished_at <= ?", finishedAt)
		countSql.Where("finished_at <= ?", finishedAt)
	}

	sql.Order("histories.started_at desc").Limit(limit).Offset(offset).Scan(&result)
	countSql.Scan(&count)

	return &result, count
}

func GetHistoryById(repositoryId string, runId string) *model.History {
	result := model.History{}
	db.Where("repository_id = ? AND run_id = ? ", repositoryId, runId).First(&result)
	return &result
}

func CreateHistory(history *model.History) *model.History {
	db.Create(history)
	return history
}

func UpdateHistory(history *model.History) *model.History {
	db.Save(history)
	return history
}

func GetRepositoryById(repositoryId string) *model.Repository {
	result := model.Repository{}
	db.Where("repository_id = ?", repositoryId).Limit(1).Find(&result)
	return &result
}

func CreateRepository(repository *model.Repository) *model.Repository {
	db.Save(repository)
	return repository
}

func GetHistryCount(repositoryName string, startedAt string, finishedAt string) (*[]custom.HistoryCounter) {
	result := []custom.HistoryCounter{}
	sql := db.Table("histories").Select("repositories.repository_name, count(1) as count").Joins("left join repositories on repositories.repository_id = histories.repository_id")

	if repositoryName != "" {
		sql.Where("repository_name LIKE ?", repositoryName + "%")
	}

	if startedAt != "" {
		sql.Where("started_at >= ?", startedAt)
	}

	if finishedAt != "" {
		sql.Where("finished_at <= ?", finishedAt)
	}

	sql.Group("repositories.repository_name").Order("count desc").Scan(&result)

	return &result
}