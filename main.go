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

	//detail := dao.UserDetail("user01")
	//fmt.Println(detail.Booking.StartTime.String(), detail.Spot.HourlyFee)
	//
	//var ti time.Time
	//fmt.Println(ti.String())

	//err := dao.UpdateLoginTime("qwe")
	//if err != nil {
	//	fmt.Println("ERROR: ", err)
	//}

	//fmt.Println(dao.GetRequiredSpot(dao.DB, "2021-01-04", "1","1","1"))

	controller.StartUp()
}
