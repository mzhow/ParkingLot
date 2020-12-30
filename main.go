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

	//detail,_ := dao.UserDetail("qwe")
	//fmt.Println(detail.Booking.StartTime.String())

	//err := dao.UpdateLoginTime("qwe")
	//if err != nil {
	//	fmt.Println("ERROR: ", err)
	//}

	controller.StartUp()
}
