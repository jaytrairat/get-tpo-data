package cfuncs

import (
	"fmt"
	"os"
	"time"

	"github.com/xuri/excelize/v2"
)

func CreateExcelFileForCaseList(excelHeaders []string, caseList [][]string, name string) error {
	currentDate := time.Now().Format("2006-01-02_15-04-05")
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
			dataCell, _ := excelize.ColumnNumberToName(j + 1)
			file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", dataCell, row), rowData)
		}
	}

	SetColumnWidths(file)
	SetStyles(file, len(caseList))

	excelFilePath := fmt.Sprintf("%s/export_at_%s_caseList.xlsx", folderName, currentDate)
	if err := file.SaveAs(excelFilePath); err != nil {
		return fmt.Errorf("ERROR :: failed to save Excel file: %w", err)
	}

	fmt.Printf("INFO :: Excel file for case list saved successfully")
	return nil
}
