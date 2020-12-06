package main

import (
	"fmt"
	"strconv"
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

	// 过滤数据，添加公司名, 按公司名分组
	baseSlice := [][][]string{}
	rows, _ := f.GetRows("蓝票")
	customerName := ""
	customerIndex := -1
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		if len(row) > 0 {
			if row[0] == "客户名称:" {
				customerName = row[1]
				customerIndex++
				baseSlice = append(baseSlice, [][]string{})
			}
			isNum := false
			for _, r := range row[0] {
				if unicode.IsNumber(r) {
					isNum = true
				}
			}
			if isNum && row[6] != "" && row[6] != "0" {
				row = append(row[:2], append([]string{customerName}, row[2:]...)...)
				baseSlice[customerIndex] = append(baseSlice[customerIndex], row)
			}
		}
	}

	// 大于limit的对半分
	for _, items := range baseSlice {
		for index, item := range items {
			itemPrice, _ := strconv.Atoi(item[7])
			if itemPrice > limit {
				itemNumber, _ := strconv.Atoi(item[5])
				splitNumber := itemNumber / 2
				fmt.Println(splitNumber)
				items = append(items[:index+1], append([][]string{item}, items[index+1:]...)...)
			}
		}
	}

	// 添加群组号
	// ticketNum := 1
	// for _, items := range baseSlice {
	// 	accrued := 0
	// 	for _, item := range items {
	// 		itemPrice, _ := strconv.Atoi(item[7])
	// 		// fmt.Println(itemPrice + accrued)
	// 		if itemPrice+accrued < limit {
	// 			accrued = itemPrice + accrued
	// 			item[0] = strconv.Itoa(ticketNum)
	// 		} else {
	// 			ticketNum++
	// 			item[0] = strconv.Itoa(ticketNum)
	// 			accrued = itemPrice
	// 		}
	// 		fmt.Println(item)
	// 	}
	// 	fmt.Println("")
	// 	ticketNum++
	// }

	for _, items := range baseSlice {
		for _, item := range items {
			fmt.Println(item)
		}
		fmt.Println("")
	}

	// if err := f.SaveAs("Book1.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }
}
