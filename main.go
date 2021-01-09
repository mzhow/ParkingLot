package main

import (
	"ParkingLot/controller"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	controller.StartUp()
}
