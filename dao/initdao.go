package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var DB *sql.DB
var DBAdmin *sql.DB

func init() {
	var err error

	// 普通用户数据库
	DB, err = sql.Open("mysql",
		"ParkingLotUser:zeVym.04wcTni.03@tcp(47.97.82.144:3306)/ParkingLot?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	// 管理员数据库
	DBAdmin, err = sql.Open("mysql",
		"ParkingLotAdmin:Qzect.03yibm@tcp(47.97.82.144:3306)/ParkingLot?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
}

func timeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
