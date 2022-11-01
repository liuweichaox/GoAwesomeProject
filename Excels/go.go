package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func main() {
	f, err := excelize.OpenFile("C:\\Users\\89370\\Desktop\\副本奖励名单（二次、发送版）(1)(21).xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	var sheet1 = "Sheet1"
	var data []ExcelModel
	for i := 2; i <= len(rows); i++ {
		var av, _ = f.GetCellValue(sheet1, "A"+strconv.Itoa(i))
		var bv, _ = f.GetCellValue(sheet1, "B"+strconv.Itoa(i))
		var cv, _ = f.GetCellValue(sheet1, "C"+strconv.Itoa(i))
		var fbv, _ = strconv.ParseFloat(bv, 64)
		data = append(data, ExcelModel{name: av, amount: fbv, phone: cv})
	}
	fmt.Println("read excel success")
	// 定义一个全局对象db
	var db *sql.DB
	// DSN:Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/dev?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	defer func() {
		db.Close()
	}()
	for _, item := range data {
		sqlStr := "insert into UserRecord(name, cust_phone,amount) values (?,?,?)"
		result, err := db.Exec(sqlStr, item.name, item.phone, item.amount)

		if err != nil {
			fmt.Printf(" error %s \n", err)
		} else {
			var lid, _ = result.LastInsertId()
			var rid, _ = result.RowsAffected()
			fmt.Printf("success LastInsertId %s  RowsAffected %s\n", lid, rid)
		}
	}
	fmt.Println("complete")
}

type ExcelModel struct {
	name   string
	amount float64
	phone  string
}
