package cfuncs

import (
	"github.com/xuri/excelize/v2"
)

var headers = []string{"ลำดับ", "File name", "SHA-256", "File size"}

func WriteCasesToExcel(cases []CaseData, filename string) error {
	f := excelize.NewFile()

	f.SetSheetRow("Sheet1", "A1", &headers)
	// headers := []string{"InstId", "TrackingCode", "StatusName"}

	// // Write headers

	// // Write data rows
	// for rowIndex, c := range cases {
	// 	rowNumber := rowIndex + 2 // Data starts on the second row
	// 	if err := f.SetCellValue(sheetName, "A"+strconv.Itoa(rowNumber), c.InstId); err != nil {
	// 		return err
	// 	}
	// 	if err := f.SetCellValue(sheetName, "B"+strconv.Itoa(rowNumber), c.TrackingCode); err != nil {
	// 		return err
	// 	}
	// 	if err := f.SetCellValue(sheetName, "C"+strconv.Itoa(rowNumber), c.StatusName); err != nil {
	// 		return err
	// 	}
	// }

	// // Set the active sheet
	// f.SetActiveSheet(index)
	f.SaveAs(filename)
	// // Save the Excel file
	// if err := f.SaveAs(filename); err != nil {
	// 	return err
	// }

	return nil
}
