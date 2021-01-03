package dao

import "ParkingLot/model"

func GetSpot(bookingId int) *model.Spot {
	query := "SELECT spot_id FROM booking WHERE booking_id=?"
	row := DB.QueryRow(query, bookingId)
	var spotId int
	row.Scan(&spotId)
	query = "SELECT spot_name, is_empty, indoor, charging, daily_fee FROM spot WHERE spot_id=?"
	row = DB.QueryRow(query, spotId)
	spot := &model.Spot{}
	row.Scan(&spot.SpotName, &spot.IsEmpty, &spot.Indoor, &spot.Charging, &spot.DailyFee)
	return spot
}

func UpdateSpot(spotId int, isEmpty int) error {
	update := "UPDATE spot SET is_empty=? WHERE spot_id=?"
	_, err := DB.Exec(update, isEmpty, spotId)
	return err
}

func GetSpotDailyFee(spotId int) (dailyFee float32) {
	query := "SELECT daily_fee FROM spot WHERE spot_id=?"
	row := DB.QueryRow(query, spotId)
	row.Scan(&dailyFee)
	return dailyFee
}
