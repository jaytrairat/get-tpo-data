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
	var limit int
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

			// cases, _ := cfuncs.GetCaseList(startDate, endDate, limit)
			// fmt.Println(len(cases))
			// if len(cases) != 0 {
			// 	for _, icase := range cases {
			// 		caseData, _ := cfuncs.GetRelatedIds(icase.InstId)
			// 		var combinedCaseData []string =
			// 	}
			// }
			var excelHeaders []string = []string{"CaseId", "CaseNumber", "NumberOfRelatedIds", "RelatedIds"}
			var mocked [][]string = [][]string{
				{"ID", "Name", "Age"}, // Header row
				{"1", "Alice", "25"},
				{"2", "Bob", "30"},
				{"3", "Charlie", "35"},
			}
			var excelName string = fmt.Sprintf("%s_%s", startDate, endDate)
			cfuncs.CreateExcelFileForCaseList(excelHeaders, mocked, excelName)

		},
	}
	rootCmd.Flags().StringVarP(&startDate, "startDate", "s", "", "Start date in YYYY-MM-DD format (default: 1 month ago)")
	rootCmd.Flags().StringVarP(&endDate, "endDate", "e", "", "End date in YYYY-MM-DD format (default: today)")
	rootCmd.Flags().IntVarP(&limit, "limit", "l", 1, "Number of rows to be extracted")
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
