package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jaytrairat/get-tpo-data/cfuncs"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	defaultEndDate := time.Now().Format("2006-01-02")
	defaultStartDate := time.Now().AddDate(0, -1, 0).Format("2006-01-02")

	var startDate string
	var endDate string
	var rootCmd = &cobra.Command{
		Use:   "get-tpo-data",
		Short: "TPO Data extractor",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if startDate == "" {
				startDate = defaultStartDate
			}
			if endDate == "" {
				endDate = defaultEndDate
			}

			cases, _ := cfuncs.GetCaseList(startDate, endDate, 1)
			fmt.Println(cases)

			casesDetail, _ := cfuncs.GetCaseDetail("881127")
			fmt.Println(casesDetail)
		},
	}
	rootCmd.Flags().StringVarP(&startDate, "startDate", "s", "", "Start date in YYYY-MM-DD format (default: 1 month ago)")
	rootCmd.Flags().StringVarP(&endDate, "endDate", "e", "", "End date in YYYY-MM-DD format (default: today)")
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
