package cfuncs

type ApiResponse struct {
	Value Value `json:"Value"`
}

type Value struct {
	Data []CaseData `json:"Data"`
}

type CaseData struct {
	InstId       int    `json:"InstId"`
	TrackingCode string `json:"TrackingCode"`
	StatusName   string `json:"StatusName"`
}
