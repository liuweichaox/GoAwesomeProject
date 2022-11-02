package main

import (
	"database/sql"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 定义一个全局对象db
	var db *sql.DB
	// DSN:Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/dev?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err := sql.Open("mysql", dsn)
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	defer func() {
		db.Close()
	}()
	if err != nil {
		fmt.Println(err)
	}

	var f = excelize.NewFile()
	var sheet1 = "Sheet1"
	index := f.NewSheet(sheet1)
	f.SetActiveSheet(index)
	f.SetCellValue(sheet1, "A1", "姓名")
	f.SetCellValue(sheet1, "B1", "手机号")
	f.SetCellValue(sheet1, "C1", "年龄")
	f.SetCellValue(sheet1, "D1", "邮箱")
	err = f.SaveAs("Book.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("complete")

}
