package custom

type HistoryCounterResponse struct {
	Counts *[]HistoryCounter `json:"counts"`
}