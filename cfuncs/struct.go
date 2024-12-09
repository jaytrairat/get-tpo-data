package cfuncs

import (
	"encoding/json"
)

type ApiResponse[T any] struct {
	Value json.RawMessage `json:"Value"`
}

type ValueWithData[T any] struct {
	Data []T `json:"Data"`
}

type CaseData struct {
	InstId       int    `json:"InstId"`
	TrackingCode string `json:"TrackingCode"`
	StatusName   string `json:"StatusName"`
}
