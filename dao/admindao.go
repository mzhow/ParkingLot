package dao

import "ParkingLot/model"


func GetAdminEncodePassword(username string) (encodePassword string) {
	db := InitDB()
	defer db.Close()
	query := "SELECT password FROM admin WHERE username='"+username+"';"
	rows, err := db.Query(query)
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&encodePassword)
		checkErr(err)
	}
	return encodePassword
}

func AdminDetail(username string) (detail model.AdminDetail){
	db := InitDB()
	defer db.Close()
	query := "SELECT user_id, role_id, username, login_time, logout_time " +
		"FROM admin " +
		"WHERE username='" + username + "' AND valid=1;"
	rows, err := db.Query(query)
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&detail.UserId, &detail.RoleId, &detail.Username,
			&detail.LoginTime, &detail.LogoutTime)
		checkErr(err)
	}
	return detail
}
