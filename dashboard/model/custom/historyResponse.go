package custom

type HistoryResponse struct {
	Count     int                  `json:"count"`
	Histories *[]HistoryRepository `json:"histories"`
}
