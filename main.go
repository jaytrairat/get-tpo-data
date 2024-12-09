package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

type ApiResponse struct {
	Data string `json:"data"`
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "get-tpo-data [url]",
		Short: "TPO Data extractor",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]
			fmt.Println(url)
			response, respError := http.Get("https://jsonplaceholder.typicode.com/todos/1")
			if respError != nil {
				fmt.Println("Error when trying to call API")
			}
			defer response.Body.Close()

			if response.StatusCode != http.StatusOK {
				log.Fatalf("API returned status code: %d", response.StatusCode)
			}

			body, ioError := io.ReadAll(response.Body)
			if ioError != nil {
				fmt.Println("Error when extract response")
			}
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
