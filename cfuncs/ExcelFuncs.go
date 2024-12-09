package cfuncs

import (
	"fmt"
	"os"
	"time"

	"github.com/xuri/excelize/v2"
)

func CreateExcelFileForCaseList(caseList []CaseData, startDate, endDate string, limit int) error {
	currentDate := time.Now().Format("2006-01-02")
	folderName := fmt.Sprintf("%s-%s_caseList", startDate, endDate)

	if err := os.MkdirAll(folderName, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create folder: %w", err)
	}

	file := excelize.NewFile()

	headers := []string{"InstId", "TrackingCode", "StatusName"}
	for col, header := range headers {
		cell, _ := excelize.ColumnNumberToName(col + 1)
		file.SetCellValue("Sheet1", fmt.Sprintf("%s1", cell), header)
	}

	for i, caseData := range caseList {
		row := i + 2
		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), caseData.InstId)
		file.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), caseData.TrackingCode)
		file.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), caseData.StatusName)
	}

	excelFilePath := fmt.Sprintf("%s/export_at_%s_limit_%d_caseList.xlsx", folderName, currentDate, limit)
	if err := file.SaveAs(excelFilePath); err != nil {
		return fmt.Errorf("failed to save Excel file: %w", err)
	}

	fmt.Printf("Excel file for case list saved successfully")
	return nil
}
