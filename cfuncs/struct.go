package cfuncs

type ApiResponse[T any] struct {
	Value Value[T] `json:"Value"`
}

type Value[T any] struct {
	Data []T `json:"Data"`
}

type CaseData struct {
	InstId       int    `json:"InstId"`
	TrackingCode string `json:"TrackingCode"`
	StatusName   string `json:"StatusName"`
}
