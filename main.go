package main

import (
	"fmt"
	"unicode"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

var (
	limit = 1130000
)

func main() {
	f, err := excelize.OpenFile("base.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	baseSlice := [][]string{}
	rows, _ := f.GetRows("蓝票")
	customerName := ""
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		if len(row) > 0 {
			if row[0] == "客户名称:" {
				customerName = row[1]
			}
			isNum := false
			for _, r := range row[0] {
				if unicode.IsNumber(r) {
					isNum = true
				}
			}
			if isNum && row[6] != "" && row[6] != "0" {
				row = append(row[:2], append([]string{customerName}, row[2:]...)...)
				baseSlice = append(baseSlice, row)
			}
		}
	}
	for _, v := range baseSlice {
		fmt.Println(v)
	}

	// rows, _ := f.GetRows("蓝票")
	// customerName := ""
	// for i := 0; i < len(rows); i++ {
	// 	row := rows[i]
	// 	if len(row) > 0 {
	// 		if row[0] == "客户名称:" {
	// 			customerName = row[1]
	// 		}
	// 		isNum := false
	// 		for _, r := range row[0] {
	// 			if unicode.IsNumber(r) {
	// 				isNum = true
	// 			}
	// 		}
	// 		if !isNum {
	// 			f.RemoveRow("蓝票", i+1)
	// 			rows, _ = f.GetRows("蓝票")
	// 			i--
	// 		} else {
	// 			f.SetCellValue("蓝票", "H"+strconv.Itoa(i+1), customerName)
	// 		}
	// 	} else {
	// 		f.RemoveRow("蓝票", i+1)
	// 		rows, _ = f.GetRows("蓝票")
	// 		i--
	// 	}
	// }
	// f.RemoveRow("蓝票", 1)
	// f.RemoveCol("蓝票", "I")
	// f.RemoveCol("蓝票", "J")

	// if err := f.SaveAs("Book1.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }
}
