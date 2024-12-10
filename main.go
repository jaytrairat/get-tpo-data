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

			fmt.Printf("Info :: Getting data from %s to %s with limit %d\n", startDate, endDate, limit)
			cases, _ := cfuncs.GetCaseList(startDate, endDate, limit)
			if len(cases) != 0 {
				fmt.Printf("Info :: %d cases found, trying to get related cases\n", len(cases))
				var excelHeaders []string = []string{"CaseId", "CaseNumber", "NumberOfRelatedIds", "RelatedIds"}
				var result [][]string
				for i, icase := range cases {
					bar := fmt.Sprintf("[%s%s]", string(cfuncs.RepeatRune('=', i)), string(cfuncs.RepeatRune(' ', len(cases)-i)))
					fmt.Printf("\rLoading... %s", bar)
					caseData, _ := cfuncs.GetRelatedIds(icase.InstId)
					if len(caseData) != 0 {
						result = append(result, []string{fmt.Sprint(icase.InstId), icase.TrackingCode, fmt.Sprint(len(caseData)), ""})
					}
				}
				fmt.Printf("\nInfo :: Select %d cases to be exported\n", len(result))

				var excelName string = fmt.Sprintf("%s_%s", startDate, endDate)
				cfuncs.CreateExcelFileForCaseList(excelHeaders, result, excelName)
			}

		},
	}
	rootCmd.Flags().StringVarP(&startDate, "startDate", "s", "", "Start date in YYYY-MM-DD format (default: 1 month ago)")
	rootCmd.Flags().StringVarP(&endDate, "endDate", "e", "", "End date in YYYY-MM-DD format (default: today)")
	rootCmd.Flags().IntVarP(&limit, "limit", "l", 1, "Number of rows to be extracted")
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
