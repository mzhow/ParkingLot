package dao

import "ParkingLot/model"

func InsertCar(carName string, isParking int) error {
	insert := "INSERT INTO car(car_name, is_parking)values(?,?)"
	_, err := DB.Exec(insert, carName, isParking)
	return err
}

func GetCarId(carName string) (carId int) {
	query := "SELECT car_id FROM car WHERE car_name=?"
	row := DB.QueryRow(query, carName)
	row.Scan(&carId)
	return carId
}

func GetCar(carId int) *model.Car {
	query := "SELECT car_name, is_parking, entry_time, out_time FROM car WHERE car_id=?"
	row := DB.QueryRow(query, carId)
	car := &model.Car{}
	row.Scan(&car.CarName, &car.IsParking, &car.EntryTime, &car.OutTime)
	return car
}

func UpdateCarEntryTime(carId int) error {
	update := "UPDATE car SET is_parking=?, entry_time=? WHERE car_id=?"
	_, err := DB.Exec(update, 1, timeNow(), carId)
	return err
}

func UpdateCarOutTime(carId int) error {
	update := "UPDATE car SET is_parking=?, out_time=? WHERE car_id=?"
	_, err := DB.Exec(update, 0, timeNow(), carId)
	return err
}
