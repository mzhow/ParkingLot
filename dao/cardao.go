package dao

func InsertCar(carName string, isParking int) {
	db := InitDB()
	defer db.Close()

	stmt, err := db.Prepare("INSERT car SET car_name=?, is_parking=?")
	checkErr(err)

	_, err = stmt.Exec(carName, isParking)
	checkErr(err)
}

func GetCarId(carName string) (carId int) {
	db := InitDB()
	defer db.Close()
	query := "SELECT car_id FROM car WHERE car_name='" + carName + "';"
	rows, err := db.Query(query)
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&carId)
		checkErr(err)
	}
	return carId
}
