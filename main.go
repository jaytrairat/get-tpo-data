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
	var startDate string
	var endDate string
	var caseId int
	var limit int
	var task string
	var rootCmd = &cobra.Command{
		Use:   "get-tpo-data",
		Short: "TPO Data extractor",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if task == "list-all" {
				totalWidth := 25
				fmt.Printf("Info :: Getting data from %s to %s with limit %d\n", startDate, endDate, limit)
				cases, _ := cfuncs.GetCaseList(startDate, endDate, limit)
				numberOfCases := len(cases.Value.Data)
				if numberOfCases != 0 {
					fmt.Printf("Info :: %d cases found, trying to get related cases\n", numberOfCases)
					var excelHeaders []string = []string{"เลขคดี", "Link", "จำนวนเคสที่เกี่ยวข้อง", "รายละเอียด", "Case ids ที่เกี่ยวข้อง"}
					var result [][]string
					for i, icase := range cases.Value.Data {
						caseDetail, _ := cfuncs.GetCaseDetailByInstId(icase.InstId)

						if len(caseDetail.Value) != 0 {
							progress := i * totalWidth / numberOfCases
							bar := fmt.Sprintf("[%s%s]",
								string(cfuncs.RepeatRune('=', progress)),
								string(cfuncs.RepeatRune(' ', totalWidth-progress)))
							fmt.Printf("\rLoading... %s %d%%", bar, progress*100/totalWidth)
							caseData, _ := cfuncs.GetRelatedCase(caseDetail.Value[0].DataId)

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
					var columnWidths = map[string]float64{
						"A": 22,
						"B": 50,
						"C": 15,
						"D": 60,
						"E": 60,
					}
					var excelName string = fmt.Sprintf("case-list_%s_%s", startDate, endDate)
					cfuncs.CreateExcelFileForCaseList(excelHeaders, result, excelName, columnWidths)
				}
			} else if task == "related-case" {
				caseDetail, _ := cfuncs.GetCaseDetailByInstId(caseId)
				if len(caseDetail.Value) != 0 {
					var excelHeaders []string = []string{"เลขรับแจ้งความ", "ประเภท", "หน่วยงานที่รับผิดชอบ", "สถานะ", "มูลค่าความเสียหาย", "รายละเอียด"}
					var result [][]string
					fmt.Println("INFO :: Getting related case")
					relatedCase, _ := cfuncs.GetRelatedCase(caseDetail.Value[0].DataId)
					if len(relatedCase.Value.Data) != 0 {
						for _, data := range relatedCase.Value.Data {
							subCase, _ := cfuncs.GetCaseDetailByCaseId(data.CaseId)
							result = append(result, []string{fmt.Sprint(data.CaseId), data.CaseType, data.OrgName, fmt.Sprint(data.CountRate), data.DamageValue, subCase.Value.CaseBehavior})
						}
					}
					var columnWidths = map[string]float64{
						"A": 30,
						"B": 50,
						"C": 50,
						"D": 50,
						"E": 30,
						"F": 80,
					}
					var excelName string = fmt.Sprintf("related-case_%d", caseId)
					cfuncs.CreateExcelFileForCaseList(excelHeaders, result, excelName, columnWidths)
				}
			} else if task == "bank-account" && caseId != 0 {
				caseDetail, _ := cfuncs.GetCaseDetailByInstId(caseId)
				if len(caseDetail.Value) != 0 {
					var excelHeaders []string = []string{"เลขบัญชี", "ชื่อบัญชี", "ธนาคาร"}
					var result [][]string
					fmt.Println("INFO :: Getting case details")
					bankList, _ := cfuncs.GetBankAccount(caseDetail.Value[0].DataId)
					if len(bankList.Value) != 0 {
						for _, bankData := range bankList.Value {
							result = append(result, []string{bankData.BankAccount, bankData.BankAccountName, bankData.BankName})
						}
					}
					var columnWidths = map[string]float64{
						"A": 30,
						"B": 50,
						"C": 50,
					}
					var excelName string = fmt.Sprintf("bank-account_%d", caseId)
					cfuncs.CreateExcelFileForCaseList(excelHeaders, result, excelName, columnWidths)
				}
			} else {
				fmt.Println("Command not found")
			}

		},
	}
	rootCmd.Flags().StringVarP(&startDate, "startDate", "s", time.Now().AddDate(0, -1, 0).Format("2006-01-02"), "Start date in YYYY-MM-DD format (default: 1 month ago)")
	rootCmd.Flags().StringVarP(&endDate, "endDate", "e", time.Now().Format("2006-01-02"), "End date in YYYY-MM-DD format (default: today)")
	rootCmd.Flags().IntVarP(&caseId, "caseId", "c", 1, "Case id from TPO")
	rootCmd.Flags().IntVarP(&limit, "limit", "l", 1, "Number of rows to be extracted")
	rootCmd.Flags().StringVarP(&task, "task", "t", "list-all", "Type of task to be executed (default:list-all) [list-all, related-case, bank-account]")
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
