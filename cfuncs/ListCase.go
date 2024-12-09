package cfuncs

import (
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

func makeGetRequest(url string) (*http.Response, error) {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+getBearerToken())

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making API request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		return nil, fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}
	return resp, nil
}

func handleApiResponse[T any](body io.Reader) ([]T, error) {
	var apiResponse ApiResponse[T]

	decoder := json.NewDecoder(body)
	if err := decoder.Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode ApiResponse: %w", err)
	}

	var valueWithData ValueWithData[T]
	if err := json.Unmarshal(apiResponse.Value, &valueWithData); err == nil {
		return valueWithData.Data, nil
	}

	var valueAsArray []T
	if err := json.Unmarshal(apiResponse.Value, &valueAsArray); err == nil {
		return valueAsArray, nil
	}

	return nil, fmt.Errorf("failed to decode Value into defined structs")
}

func GetCaseList(startDate, endDate string, limit int) ([]CaseData, error) {
	const listCasesAPIURL = "https://officer.thaipoliceonline.go.th/api/e-form/v1.0/BpmProcInst/workflow/task-list-new?RequireTotalCount=true&Ext2=3527&RoleCode=MNG_BKK&Offset=0&Length=%d&SortDesc=true&StartDate=%s&EndDate=%s&CategoryId=1&RequireStuckCase=false"

	url := fmt.Sprintf(listCasesAPIURL, limit, startDate, endDate)
	response, err := makeGetRequest(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return handleApiResponse[CaseData](response.Body)
}

func GetCaseDetail(caseId string) ([]CaseDetailData, error) {
	const caseDetailAPIURL = "https://officer.thaipoliceonline.go.th/api/e-form/v1.0/BpmProcInstLog?instId=%s&excludeSystemCreate=true"

	url := fmt.Sprintf(caseDetailAPIURL, caseId)
	response, err := makeGetRequest(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return handleApiResponse[CaseDetailData](response.Body)
}
