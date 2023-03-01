package custom

import (
  "time"
)

type HistoryRepository struct {
  RepositoryId   string `gorm:"primaryKey;column:repository_id" json:"repository_id"`
  RepositoryName string    `gorm:"column:repository_name" json:"repository_name"`
  RunId          string `gorm:"primaryKey;column:run_id" json:"run_id"`
  WorkflowRef    string `gorm:"column:workflow_ref" json:"workflow_ref"`
  JobName        string `gorm:"column:job_name" json:"job_name"`
  Status         string `gorm:"column:status" json:"status"`
  StartedAt      time.Time `gorm:"column:started_at" json:"started_at"`
  FinishedAt     time.Time `gorm:"column:finished_at" json:"finished_at"`
}