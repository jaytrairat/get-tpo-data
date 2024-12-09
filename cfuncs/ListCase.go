package cfuncs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func ListCase() {
	LIST_CASES_API := "https://officer.thaipoliceonline.go.th/api/e-form/v1.0/BpmProcInst/workflow/task-list-new?RequireTotalCount=true&Ext2=3527&RoleCode=MNG_BKK&Offset=0&Length=1&SortDesc=true&StartDate=%s&EndDate=%s&CategoryId=1&RequireStuckCase=false"

	bearerToken := os.Getenv("BEARER_TOKEN")

	if bearerToken == "" {
		log.Fatalf("BEARER_TOKEN not set in environment variables")
	}

	formattedUrl := fmt.Sprintf(LIST_CASES_API, "2024-12-01", "2024-12-05")
	fmt.Println(formattedUrl)
	req, err := http.NewRequest("GET", formattedUrl, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+bearerToken)

	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making API request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("API returned status code: %d", response.StatusCode)
	}

	var apiResponse ApiResponse
	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		log.Fatalf("Error decoding JSON response: %v", err)
	}

	// Print the response (everything as string)
	fmt.Printf("Extracted Data: %+v\n", apiResponse)

}
