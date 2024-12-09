package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

type ApiResponse struct {
	Data string `json:"data"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var rootCmd = &cobra.Command{
		Use:   "get-tpo-data",
		Short: "TPO Data extractor",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			apiUrl := os.Getenv("API_URL")
			bearerToken := os.Getenv("BEARER_TOKEN")

			if apiUrl == "" || bearerToken == "" {
				log.Fatalf("API_URL or BEARER_TOKEN not set in environment variables")
			}

			req, err := http.NewRequest("GET", apiUrl, nil)
			if err != nil {
				log.Fatalf("Error creating request: %v", err)
			}

			req.Header.Set("Authorization", "Bearer "+bearerToken)

			client := &http.Client{}
			response, err := client.Do(req)
			if err != nil {
				log.Fatalf("Error making API request: %v", err)
			}
			defer response.Body.Close()

			if response.StatusCode != http.StatusOK {
				log.Fatalf("API returned status code: %d", response.StatusCode)
			}

			body, err := io.ReadAll(response.Body)
			if err != nil {
				log.Fatalf("Error reading response body: %v", err)
			}

			fmt.Println("Response Body (raw):")
			fmt.Println(string(body))

			var apiResponse ApiResponse
			if err := json.Unmarshal(body, &apiResponse); err != nil {
				log.Fatalf("Error unmarshalling JSON: %v", err)
			}

			fmt.Printf("Extracted Data: %s\n", apiResponse.Data)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
