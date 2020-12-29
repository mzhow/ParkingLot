package dao

import "ParkingLot/model"

func GetEncodePassword(username string) (encodePassword string) {

	//sql := "select id,username,password,email from users where username = ? and password = ?"
	//row := utils.Db.QueryRow(sql, username, password)
	//user := &model.User{}
	//row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	//return user, nil
	query := "SELECT password FROM user WHERE username=?"
	row := DB.QueryRow(query, username)
	row.Scan(&encodePassword)
	return encodePassword
}

func UserDetail(username string) (detail *model.UserDetail) {
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

func CheckUsernameValid(username string) bool {
	db := InitDB()
	defer db.Close()
	query := "SELECT count(*) FROM user WHERE username='" + username +
		"' AND valid=1;"
	rows, err := db.Query(query)
	checkErr(err)
	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		checkErr(err)
	}
	return count > 0
}

func ExistUsernameOrCarName(username string, carName string) bool {
	db := InitDB()
	defer db.Close()
	que := db.QueryRow()
	query := "SELECT COUNT(*) FROM user WHERE username='" + username + "';"
	rows, err := db.Query(query)
	checkErr(err)
	var countUser int
	for rows.Next() {
		err = rows.Scan(&countUser)
		checkErr(err)
	}
	query = "SELECT COUNT(*) FROM car WHERE car_name='" + carName + "';"
	rows, err = db.Query(query)
	checkErr(err)
	var countCarName int
	for rows.Next() {
		err = rows.Scan(&countCarName)
		checkErr(err)
	}
	if countUser > 0 || countCarName > 0 {
		return true
	} else {
		return false
	}
}

func InsertUser(roleId int, carId int, username string, password string, valid int) {
	db := InitDB()
	defer db.Close()

	stmt, err := db.Prepare("INSERT user SET role_id=?, car_id=?, username=?, password=?, valid=?")
	checkErr(err)

	_, err = stmt.Exec(roleId, carId, username, password, valid)
	checkErr(err)
}
