package dao

import "ParkingLot/model"

func GetEncodePassword(username string) (encodePassword string) {
	db := InitDB()
	defer db.Close()
	query := "SELECT password FROM user WHERE username='"+username+"';"
	rows, err := db.Query(query)
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&encodePassword)
		checkErr(err)
	}
	return encodePassword
}

func UserDetail(username string) (detail *model.UserDetail){
	db := InitDB()
	defer db.Close()
	query := "SELECT u.user_id, u.role_id, u.username, " +
		"c.car_name, c.is_parking, c.entry_time, c.out_time, " +
		"s.spot_name, s.indoor, s.charging, s.hourly_fee, " +
		"b.start_time, b.end_time " +
		"FROM user u " +
		"INNER JOIN car c ON u.car_id=c.car_id " +
		"INNER JOIN spot s ON c.spot_id=s.spot_id " +
		"INNER JOIN booking b ON u.booking_id=b.booking_id " +
		"WHERE u.username='" + username + "' AND u.valid=1;"
	rows, err := db.Query(query)
	checkErr(err)
	for rows.Next() {
		detail := model.UserDetail{}
		err = rows.Scan(&detail.UserId, &detail.RoleId, &detail.Username,
			&detail.Car.CarName, &detail.Car.IsParking, &detail.Car.EntryTime, &detail.Car.OutTime,
			&detail.Spot.SpotName, &detail.Spot.Indoor, &detail.Spot.Charging, &detail.Spot.HourlyFee,
			&detail.Booking.StartTime, &detail.Booking.EndTime)
		checkErr(err)
	}
	return detail
}

