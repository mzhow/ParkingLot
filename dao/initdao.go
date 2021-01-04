package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql",
		"goweb:123456@tcp(47.97.82.144:3306)/ParkingLot?charset=utf8&parseTime=true&loc=Local")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func timeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
