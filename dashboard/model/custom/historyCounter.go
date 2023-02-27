package custom

type HistoryCounter struct {
  RepositoryName string    `gorm:"column:repository_name" json:"repository_name"`
	Count int `gorm:"column:count" json:"count"`
}