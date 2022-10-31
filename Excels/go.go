package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	CreateExcel()
}

func CreateExcel() {
	f := excelize.NewFile()
	// 創建一個工作表
	index := f.NewSheet("Sheet2")
	// 設置存儲格的值
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// 設置活頁簿的默認工作表
	f.SetActiveSheet(index)
	// 根據指定路徑保存活頁簿
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
