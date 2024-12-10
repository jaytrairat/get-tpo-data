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

func GetCaseList(startDate, endDate string, limit int) (StCaseList, error) {
	const apiUrl = "https://officer.thaipoliceonline.go.th/api/e-form/v1.0/BpmProcInst/workflow/task-list-new?RequireTotalCount=true&StatusCode=&StateCode=&ProcessedStateCode=&RoleCode=ROOT&Offset=0&Length=%d&SortSelector=TrackingCode&SortDesc=true&StartDate=%s&EndDate=%s&CategoryId=1&Casetype=&IsCheck=&RequireStuckCase=false"

	url := fmt.Sprintf(apiUrl, limit, startDate, endDate)
	response, err := makeRequest(url, "GET", nil)
	if err != nil {
		return StCaseList{}, err
	}
	defer response.Body.Close()

	var caseList StCaseList
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&caseList); err != nil {
		return StCaseList{}, fmt.Errorf("failed to decode caseList api: %w", err)
	}

	return caseList, nil
}

func GetRelatedIds(caseId int) (StRelatedCase, error) {
	const apiUrl = "https://officer.thaipoliceonline.go.th/api/ccib/v1.0/CmsOnlineCaseInfo/%d/relation"

	url := fmt.Sprintf(apiUrl, caseId)

	data, _ := json.Marshal(map[string]string{
		"Offset": "0",
		"Length": "10000",
	})

	response, err := makeRequest(url, "POST", bytes.NewReader(data))
	if err != nil {
		return StRelatedCase{}, err
	}

	defer response.Body.Close()

	var relatedCase StRelatedCase
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&relatedCase); err != nil {
		return StRelatedCase{}, fmt.Errorf("failed to decode relatedCase api: %w", err)
	}

	return relatedCase, nil
}

func GetCaseDetail(caseId int) (StCaseDetail, error) {
	const apiUrl = "https://officer.thaipoliceonline.go.th/api/e-form/v1.0/BpmProcInstLog?instId=%d&excludeSystemCreate=true"
	url := fmt.Sprintf(apiUrl, caseId)
	response, err := makeRequest(url, "GET", nil)
	if err != nil {
		return StCaseDetail{}, err
	}

	defer response.Body.Close()
	var caseDetail StCaseDetail
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&caseDetail); err != nil {
		return StCaseDetail{}, fmt.Errorf("failed to decode relatedCase api: %w", err)
	}

	return caseDetail, nil
}
