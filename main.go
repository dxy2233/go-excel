package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("base.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	newF := excelize.NewFile()
	index := newF.NewSheet("模板")
	newF.SetActiveSheet(index)

	rows, _ := f.GetRows("蓝票")
	for rowIndex, row := range rows {
		for colIndex, colCell := range row {
			if colIndex > 0 {
				break
			}
			for _, r := range colCell {
				if unicode.IsNumber(r) {
					newF.SetSheetRow("模板", "A"+strconv.Itoa(rowIndex), &row)
				}
			}
		}
	}

	// f.RemoveRow("蓝票", 1)
	// f.DuplicateRowTo("蓝票", 1, 2)
	// rows, _ = f.GetRows("蓝票")
	// for _, row := range rows {
	// 	fmt.Println(row[0])
	// }

	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
