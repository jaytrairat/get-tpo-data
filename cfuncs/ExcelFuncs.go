package cfuncs

import (
	"fmt"
	"os"
	"time"

	"github.com/xuri/excelize/v2"
)

func numberToExcelColumn(n int) string {
	if n <= 0 {
		return ""
	}
	column := ""
	for n > 0 {
		n--
		remainder := n % 26
		column = string('A'+remainder) + column
		n = n / 26
	}
	return column
}
func CreateExcelFileForCaseList(excelHeaders []string, caseList [][]string, name string) error {
	currentDate := time.Now().Format("2006-01-02")
	folderName := fmt.Sprintf("%s_caseList", name)

	if err := os.MkdirAll(folderName, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create folder: %w", err)
	}

	file := excelize.NewFile()

	for col, header := range excelHeaders {
		cell, _ := excelize.ColumnNumberToName(col + 1)
		file.SetCellValue("Sheet1", fmt.Sprintf("%s1", cell), header)
	}

	for i, caseData := range caseList {
		row := i + 2
		for j, rowData := range caseData {
			file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", numberToExcelColumn(j), row), rowData)
		}
	}

	excelFilePath := fmt.Sprintf("%s/export_at_%s_caseList.xlsx", folderName, currentDate)
	if err := file.SaveAs(excelFilePath); err != nil {
		return fmt.Errorf("failed to save Excel file: %w", err)
	}

	fmt.Printf("Excel file for case list saved successfully")
	return nil
}
