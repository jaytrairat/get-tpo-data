package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jaytrairat/get-tpo-data/cfuncs"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var rootCmd = &cobra.Command{
		Use:   "get-tpo-data",
		Short: "TPO Data extractor",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			listCasesApiUrl := os.Getenv("LIST_CASES_API")
			bearerToken := os.Getenv("BEARER_TOKEN")

			if listCasesApiUrl == "" || bearerToken == "" {
				log.Fatalf("LIST_CASES_API or BEARER_TOKEN not set in environment variables")
			}

			formattedUrl := fmt.Sprintf(listCasesApiUrl, "2024-12-01", "2024-12-05")
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

			var apiResponse cfuncs.ApiResponse
			if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
				log.Fatalf("Error decoding JSON response: %v", err)
			}

			// Print the response (everything as string)
			fmt.Printf("Extracted Data: %+v\n", apiResponse)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
