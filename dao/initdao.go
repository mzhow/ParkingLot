package dao

import (
	"database/sql"
	"fmt"
	"os"
)

func InitDB() *sql.DB{
	db, err := sql.Open("mysql",
		"goweb:123456@tcp(47.97.82.144:3306)/ParkingLot?charset=utf8&parseTime=true&loc=Local")
	checkErr(err)
	return db
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
