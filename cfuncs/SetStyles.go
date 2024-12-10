package cfuncs

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

const (
	startDataRow = 2
	startColumn  = 1
)

func SetColumnWidths(f *excelize.File, columnWidths map[string]float64) error {
	for col, width := range columnWidths {
		if err := f.SetColWidth("Sheet1", col, col, width); err != nil {
			return fmt.Errorf("error setting column width: %w", err)
		}
	}
	return nil
}

func SetStyles(f *excelize.File, recordCount int, columnCount int) error {
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:   true,
			Size:   16,
			Family: "TH Sarabun New",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	cell, _ := excelize.ColumnNumberToName(columnCount + 1)
	f.SetCellStyle("Sheet1", "A1", fmt.Sprintf("%s1", cell), headerStyle)

	indexStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Size:   16,
			Family: "TH Sarabun New",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})

	for i := startDataRow; i <= recordCount+1; i++ {
		for j := startColumn; j <= columnCount+1; j++ {
			cell, _ := excelize.ColumnNumberToName(j + 1)
			f.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", cell, i), fmt.Sprintf("%s%d", cell, i), indexStyle)
		}
	}

	return nil
}
