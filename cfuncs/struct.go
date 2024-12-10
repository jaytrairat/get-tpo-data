package cfuncs

import (
	"encoding/json"
)

type ApiResponse[T any] struct {
	Value json.RawMessage `json:"Value"`
}

type ValueWithData[T any] struct {
	Data       []T `json:"Data"`
	TotalCount int `json:"TotalCount"`
}

type CaseData struct {
	InstId       int    `json:"InstId"`
	TrackingCode string `json:"TrackingCode"`
	StatusName   string `json:"StatusName"`
	DataId       string `json:"DATA_ID"`
}

type RelatedCase struct {
	CaseId   int    `json:"CASE_ID"`
	InstId   int    `json:"INST_ID"`
	CaseNo   string `json:"CASE_NO"`
	CaseType string `json:"CASE_TYPE_ABBR"`
}
