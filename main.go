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

func ticketNumToThree(num int) string {
	s := strconv.Itoa(num)
	if len(s) == 1 {
		s = "00" + s
	}
	if len(s) == 2 {
		s = "0" + s
	}
	return s
}

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
			row = row[:7]
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
			itemSum, _ := strconv.Atoi(item[7])
			if itemSum > limit {
				// first
				itemNumber, _ := strconv.Atoi(item[5])
				splitNumber := itemNumber / 2
				item[5] = strconv.Itoa(splitNumber)
				itemPrice, _ := strconv.Atoi(item[6])
				item[7] = strconv.Itoa(splitNumber * itemPrice)
				// second
				newItem := make([]string, len(item), cap(item))
				copy(newItem, item)
				newItem[5] = strconv.Itoa(itemNumber - splitNumber)
				newItem[7] = strconv.Itoa((itemNumber - splitNumber) * itemPrice)
				items = append(items[:index+1], append([][]string{newItem}, items[index+1:]...)...)
			}
		}
	}

	// 添加群组号
	ticketNum := 1
	for _, items := range baseSlice {
		accrued := 0
		for _, item := range items {
			itemPrice, _ := strconv.Atoi(item[7])
			if itemPrice+accrued < limit {
				accrued = itemPrice + accrued
				item[0] = ticketNumToThree(ticketNum)
			} else {
				ticketNum++
				item[0] = ticketNumToThree(ticketNum)
				accrued = itemPrice
			}
		}
		ticketNum++
	}

	// 生成excel
	res := excelize.NewFile()
	sheetName := res.NewSheet("模板")
	res.SetActiveSheet(sheetName)
	// 表头
	res.SetSheetRow("模板", "A1", &[]interface{}{
		"群组", "发票明细", "客户", "单位", "规格型号", "数量", "单价", "金额",
		"备注", 1, 2, 3, 4, 5, 6, 7, "收款人", "审核人", "默认税率", "开票类型", "折扣金额"})
	index := 2
	for _, items := range baseSlice {
		for _, item := range items {
			res.SetSheetRow("模板", "A"+strconv.Itoa(index), &item)
			index++
		}
	}
	// 格式转换
	cols, _ := res.GetCols("模板")
	for colIndex, col := range cols {
		for rowIndex, rowCell := range col {
			if colIndex == 5 || colIndex == 6 || colIndex == 7 {
				if rowIndex > 0 {
					val, _ := strconv.Atoi(rowCell)
					cellCoordinates, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+1)
					res.SetCellInt("模板", cellCoordinates, val)
				}
			}
		}
	}
	if err := res.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
