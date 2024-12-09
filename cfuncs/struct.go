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

type CaseDetailData struct {
	InstId           int    `json:"INST_ID"`
	TrackingCode     string `json:"TRACKING_CODE"`
	Remark           string `json:"REMARK"`
	PersonalFullName string `json:"PERSONAL_FULL_NAME"`
}
