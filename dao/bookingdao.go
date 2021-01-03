package dao

import "ParkingLot/model"

func GetBookingTime(bookingId int) *model.Booking {
	query := "SELECT start_time, end_time FROM booking WHERE booking_id=?"
	row := DB.QueryRow(query, bookingId)
	booking := &model.Booking{}
	row.Scan(&booking.StartTime, &booking.EndTime)
	return booking
}

func GetBookingCarAndSpot(bookingId int) (carId int, spotId int) {
	query := "SELECT car_id, spot_id FROM booking WHERE booking_id=?"
	row := DB.QueryRow(query, bookingId)
	row.Scan(&carId, &spotId)
	return carId, spotId
}

func UpdateBookingValid(bookingId int, valid int) error {
	update := "UPDATE booking SET valid=? WHERE booking_id=?"
	_, err := DB.Exec(update, valid, bookingId)
	return err
}
