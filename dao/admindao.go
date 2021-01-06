package dao

import (
	"ParkingLot/model"
	"strconv"
)

func GetAdminEncodePassword(username string) (encodePassword string, err error) {
	query := "SELECT password FROM admin WHERE username=?"
	row := DBAdmin.QueryRow(query, username)
	err = row.Scan(&encodePassword)
	return encodePassword, err
}

func CheckAdminRole(username string, roleId int) bool {
	query := "SELECT count(*) FROM admin WHERE username=? AND role_id=?"
	row := DBAdmin.QueryRow(query, username, roleId)
	var count int
	row.Scan(&count)
	return count > 0
}

func CountUsers() int {
	query := "SELECT count(*) FROM user"
	row := DB.QueryRow(query)
	var count int
	row.Scan(&count)
	return count
}

func GetUsers(page string) ([]*model.User, error) {
	pagenum, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}
	pageOffset := strconv.Itoa((pagenum - 1) * defaultPageSize)
	pageSize := strconv.Itoa(defaultPageSize)
	query := "SELECT user_id, role_id, car_id, booking_id, username, login_time, logout_time, fee, valid FROM user " +
		"ORDER BY user_id LIMIT " + pageSize + " OFFSET " + pageOffset + ";"
	rows, err := DBAdmin.Query(query)
	if err != nil {
		return nil, err
	}
	var users []*model.User
	for rows.Next() {
		user := &model.User{}
		err = rows.Scan(&user.UserId, &user.RoleId, &user.CarId, &user.BookingId, &user.Username,
			&user.LoginTime, &user.LogoutTime, &user.Fee, &user.Valid)
		users = append(users, user)
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}

// 更新用户登录、登出时间
func UpdateAdminLoginTime(username string) error {
	update := "UPDATE admin SET login_time=? WHERE username=?"
	_, err := DB.Exec(update, timeNow(), username)
	return err
}
func UpdateAdminLogoutTime(username string) error {
	update := "UPDATE admin SET logout_time=? WHERE username=?"
	_, err := DB.Exec(update, timeNow(), username)
	return err
}

func GetSpots(page string) ([]*model.SpotDetail, error) {
	pagenum, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}
	pageOffset := strconv.Itoa((pagenum - 1) * defaultPageSize)
	pageSize := strconv.Itoa(defaultPageSize)
	query := "SELECT spot_id, spot_name, is_empty, indoor, charging, daily_fee, valid FROM spot " +
		"ORDER BY spot_id LIMIT " + pageSize + " OFFSET " + pageOffset + ";"
	rows, err := DBAdmin.Query(query)
	if err != nil {
		return nil, err
	}
	var spots []*model.SpotDetail
	for rows.Next() {
		spot := &model.SpotDetail{}
		err = rows.Scan(&spot.SpotId, &spot.SpotName, &spot.IsEmpty, &spot.Indoor, &spot.Charging, &spot.DailyFee, &spot.Valid)
		spots = append(spots, spot)
		if err != nil {
			return nil, err
		}
	}
	return spots, nil
}

func CountSpots() int {
	query := "SELECT count(*) FROM spot"
	row := DB.QueryRow(query)
	var count int
	row.Scan(&count)
	return count
}

func GetBookings(page string) ([]*model.BookingDetail, error) {
	pagenum, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}
	pageOffset := strconv.Itoa((pagenum - 1) * defaultPageSize)
	pageSize := strconv.Itoa(defaultPageSize)
	query := "SELECT booking_id, car_id, spot_id, start_time, end_time, valid FROM booking " +
		"ORDER BY booking_id LIMIT " + pageSize + " OFFSET " + pageOffset + ";"
	rows, err := DBAdmin.Query(query)
	if err != nil {
		return nil, err
	}
	var bookings []*model.BookingDetail
	for rows.Next() {
		booking := &model.BookingDetail{}
		err = rows.Scan(&booking.BookingId, &booking.CarId, &booking.SpotId, &booking.StartTime, &booking.EndTime, &booking.Valid)
		bookings = append(bookings, booking)
		if err != nil {
			return nil, err
		}
	}
	return bookings, nil
}

func CountBookings() int {
	query := "SELECT count(*) FROM booking"
	row := DB.QueryRow(query)
	var count int
	row.Scan(&count)
	return count
}
