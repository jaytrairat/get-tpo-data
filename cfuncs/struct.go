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

type StCaseList struct {
	Value struct {
		Data []struct {
			InstId       int    `json:"InstId"`
			TrackingCode string `json:"TrackingCode"`
			OptionalData string `json:"OptionalData"`
		} `json:"Data"`
	} `json:"Value"`
}

type StRelatedCase struct {
	Value struct {
		Data []struct {
			CaseId   int    `json:"CASE_ID"`
			InstId   int    `json:"INST_ID"`
			CaseNo   string `json:"CASE_NO"`
			CaseType string `json:"CASE_TYPE_ABBR"`
		} `json:"Data"`
	} `json:"Value"`
}

type StCaseDetail struct {
	Value []struct {
		DataId int `json:"DATA_ID"`
	} `json:"Value"`
}

type StBankAccount struct {
	Value []struct {
		BankAccount     string `json:"BANK_ACCOUNT"`
		BankAccountName string `json:"BANK_ACCOUNT_NAME"`
		BankName        string `json:"BANK_NAME"`
	} `json:"Value"`
}
