package excel

import (
	"encoding/json"
	"fmt"

	"github.com/xuri/excelize/v2"
)

const (
	
)

type ExcelFile struct {
	file *excelize.File
	name string
}

// ======================= File-level Operations ===========================
// CreateNewFile initializes a new Excel file with the given name
func CreateNewFile(name string) *ExcelFile {
	return &ExcelFile{
		file: excelize.NewFile(),
		name: name,
	}
}

// SaveAs saves the Excel file to the specified path.
func (ef *ExcelFile) SaveAs(path string) {
	if path[len(path)-1] != '/' {
		path = path + "/"
	}
	if err := ef.file.SaveAs(path + ef.name + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}

func (ef *ExcelFile) Close() error {
	return ef.file.Close()
}

// ======================= Sheet-level Operations ===========================
// CreateNewSheet adds a new sheet to the Excel file with the specified name.
func (ef *ExcelFile) CreateNewSheet(sheetName string) {
	ef.file.NewSheet(sheetName)
}

func (ef *ExcelFile) DeleteSheet(sheetName string) {
	ef.file.DeleteSheet(sheetName)
}

// Setting the sheet to be the active one
func (ef *ExcelFile) Activatesheet(sheetName string) {
	if index, err := ef.file.GetSheetIndex(sheetName); err != nil {
		fmt.Println(err)
	} else {
		ef.file.SetActiveSheet(index)
	}

}

// ======================= Row-level Operations ===========================
// Handling row level operations
func (ef *ExcelFile) SetRowValues(sheetName string, rowIndex int, values []interface{}) {
	axis := fmt.Sprintf("A%d", rowIndex)
	ef.file.SetSheetRow(sheetName, axis, &values)
}

// ======================= Cell-level Operations ===========================
// Handling cell level operations
func (ef *ExcelFile) SetCellValue(sheetName string, axis string, value interface{}) {
	ef.file.SetCellValue(sheetName, axis, &value)
}

func (ef *ExcelFile) GetCellValue(sheetName string, axis string) (string, error) {
	return ef.file.GetCellValue(sheetName, axis)
}

// ======================= Column-level Operations ===========================
func (ef *ExcelFile) SetColWidth(sheetName string, startCol, endCol string, width float64) error {
	return ef.file.SetColWidth(sheetName, startCol, endCol, width)
}

// ======================= Style-level Operations ===========================
// SetCellStyle applies a style to a range of cells in the specified sheet.
func (ef *ExcelFile) SetCellStyle(sheetName string, startCell, endCell string, styleID int) error {
	return ef.file.SetCellStyle(sheetName, startCell, endCell, styleID)
}

// SetCellStyleFromJSON applies a style to a range of cells in a sheet based on a JSON description of the style.
func (ef *ExcelFile) SetCellStyleFromJSON(sheetName string, startCell, endCell string, styleJSON string) error {
	var style excelize.Style
	if err := json.Unmarshal([]byte(styleJSON), &style); err != nil {
		return fmt.Errorf("error unmarshaling style JSON: %v", err)
	}

	// Create a new style in the file.
	styleID, err := ef.file.NewStyle(&style)
	if err != nil {
		return fmt.Errorf("error creating new style: %v", err)
	}

	err = ef.SetCellStyle(sheetName, startCell, endCell, styleID)
	if err != nil {
		return fmt.Errorf("error setting new style: %v", err)
	}
	return nil
}
