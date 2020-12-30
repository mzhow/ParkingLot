package dao

import (
	"ParkingLot/model"
)

func GetEncodePassword(username string) (string, error) {
	query := "SELECT password FROM user WHERE username=?"
	row := DB.QueryRow(query, username)
	var encodePassword string
	err := row.Scan(&encodePassword)
	return encodePassword, err
}

func UserDetail(username string) *model.UserDetail {
	//query := "SELECT u.user_id, u.role_id, u.username, " +
	//	"c.car_name, c.is_parking, c.entry_time, c.out_time, " +
	//	"s.spot_name, s.is_empty, s.indoor, s.charging, s.hourly_fee, " +
	//	"c.car_name, s.spot_name, b.start_time, b.end_time " +
	//	"FROM user u " +
	//	"INNER JOIN car c ON u.car_id=c.car_id " +
	//	"INNER JOIN spot s ON c.spot_id=s.spot_id " +
	//	"INNER JOIN booking b ON u.booking_id=b.booking_id " +
	//	"WHERE u.username=? AND u.valid=1;"
	//row := DB.QueryRow(query, username)
	//detail := &model.UserDetail{}
	//row.Scan(&detail.UserId, &detail.RoleId, &detail.Username,
	//		&detail.Car.CarName, &detail.Car.IsParking, &detail.Car.EntryTime, &detail.Car.OutTime,
	//		&detail.Spot.SpotName, &detail.Spot.IsEmpty, &detail.Spot.Indoor, &detail.Spot.Charging, &detail.Spot.HourlyFee,
	//		&detail.Booking.CarName, &detail.Booking.SpotName, &detail.Booking.StartTime, &detail.Booking.EndTime)
	//return detail

	query := "SELECT user_id, role_id, username, car_id, booking_id FROM user WHERE valid=1 AND username=?"
	row := DB.QueryRow(query, username)
	detail := &model.UserDetail{}
	var carId, bookingId int
	row.Scan(&detail.UserId, &detail.RoleId, &detail.Username, &carId, &bookingId)
	detail.Car = GetCar(carId)
	detail.Booking = GetBooking(bookingId)
	return detail
}

func CheckUsernameValid(username string) bool {
	query := "SELECT count(*) FROM user WHERE username=? AND valid=1"
	row := DB.QueryRow(query, username)
	var count int
	row.Scan(&count)
	return count > 0
}

func ExistUsernameOrCarName(username string, carName string) bool {
	query := "SELECT COUNT(*) FROM user WHERE username=?"
	row := DB.QueryRow(query, username)
	var countUser int
	row.Scan(&countUser)
	query = "SELECT COUNT(*) FROM car WHERE car_name=?"
	row = DB.QueryRow(query, carName)
	var countCarName int
	row.Scan(&countCarName)
	if countUser > 0 || countCarName > 0 {
		return true
	} else {
		return false
	}
}

func InsertUser(roleId int, carId int, username string, password string, valid int) error {
	insert := "INSERT INTO user(role_id,car_id,username,password,valid)values(?,?,?,?,?)"
	_, err := DB.Exec(insert, roleId, carId, username, password, valid)
	return err
}

func UpdateLoginTime(username string) error {
	update := "UPDATE user SET login_time=? WHERE username=?"
	_, err := DB.Exec(update, timeNow(), username)
	return err
}
