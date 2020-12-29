package main

import (
	"ParkingLot/controller"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//password := "12345678"
	//encode := controller.HashAndSalt(password)
	//fmt.Println(encode)
	//if controller.ComparePasswords(encode, password) {
	//	fmt.Println("true")
	//}
	controller.StartUp()
}
