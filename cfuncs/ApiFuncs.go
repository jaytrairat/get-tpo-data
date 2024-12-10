package cfuncs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func getBearerToken() string {
	bearerToken := os.Getenv("BEARER_TOKEN")
	if bearerToken == "" {
		log.Fatalf("BEARER_TOKEN not set in environment variables")
	}
	return bearerToken
}

func makeRequest(url, method string, body io.Reader) (*http.Response, error) {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+getBearerToken())
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making API request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	return resp, nil
}

func DecodeApiResponse[TargetStruct any](body io.Reader) ([]TargetStruct, error) {
	var apiResponse ApiResponse[TargetStruct]

	decoder := json.NewDecoder(body)
	if err := decoder.Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode ApiResponse: %w", err)
	}

	var valueWithData ValueWithData[TargetStruct]
	if err := json.Unmarshal(apiResponse.Value, &valueWithData); err == nil {
		return valueWithData.Data, nil
	}

	return nil, fmt.Errorf("failed to decode Value into defined structs")
}

func GetCaseList(startDate, endDate string, limit int) ([]StCaseList, error) {
	const listCasesAPIURL = "https://officer.thaipoliceonline.go.th/api/e-form/v1.0/BpmProcInst/workflow/task-list-new?StatusCode=&StateCode=&ProcessedStateCode=&Ext2=3527&RoleCode=MNG_BKK&Offset=1&Length=%d&SortSelector=TrackingCode&SortDesc=true&StartDate=%s&EndDate=%s&CategoryId=1&Casetype=&IsCheck=&RequireStuckCase=false"

	url := fmt.Sprintf(listCasesAPIURL, limit, startDate, endDate)
	response, err := makeRequest(url, "GET", nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return DecodeApiResponse[StCaseList](response.Body)
}

func GetRelatedIds(caseId int) ([]StRelatedCase, error) {
	const listCasesAPIURL = "https://officer.thaipoliceonline.go.th/api/ccib/v1.0/CmsOnlineCaseInfo/%d/relation"

	url := fmt.Sprintf(listCasesAPIURL, caseId)
	data, _ := json.Marshal(map[string]string{
		"Offset": "0",
		"Length": "10000",
	})

	response, err := makeRequest(url, "POST", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return DecodeApiResponse[StRelatedCase](response.Body)
}
