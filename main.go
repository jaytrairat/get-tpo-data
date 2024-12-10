package main

import (
	"fmt"
	"log"
	"strings"
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
			totalWidth := 25
			fmt.Printf("Info :: Getting data from %s to %s with limit %d\n", startDate, endDate, limit)
			cases, _ := cfuncs.GetCaseList(startDate, endDate, limit)
			numberOfCases := len(cases.Value.Data)
			if numberOfCases != 0 {
				fmt.Printf("Info :: %d cases found, trying to get related cases\n", numberOfCases)
				var excelHeaders []string = []string{"เลขคดี", "Link", "จำนวนเคสที่เกี่ยวข้อง", "รายละเอียด", "Case ids ที่เกี่ยวข้อง"}
				var result [][]string
				for i, icase := range cases.Value.Data {
					caseDetail, _ := cfuncs.GetCaseDetail(icase.InstId)

					if len(caseDetail.Value) != 0 {
						progress := i * totalWidth / numberOfCases
						bar := fmt.Sprintf("[%s%s]",
							string(cfuncs.RepeatRune('=', progress)),
							string(cfuncs.RepeatRune(' ', totalWidth-progress)))
						fmt.Printf("\rLoading... %s %d%%", bar, progress*100/totalWidth)
						caseData, _ := cfuncs.GetRelatedIds(caseDetail.Value[0].DataId)

						if len(caseData.Value.Data) != 0 {
							var caseNos []string

							for _, item := range caseData.Value.Data {
								caseNos = append(caseNos, item.CaseNo)
							}
							result = append(result, []string{icase.TrackingCode, fmt.Sprintf("https://officer.thaipoliceonline.go.th/pct-in/officer/task-admin-view/%d#task-admin", icase.InstId), fmt.Sprint(len(caseData.Value.Data)), icase.OptionalData, strings.Join(caseNos, ",")})
						}
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
