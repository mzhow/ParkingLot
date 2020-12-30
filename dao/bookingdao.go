package dao

import "ParkingLot/model"

func GetBooking(bookingId int) *model.Booking {
	query := "SELECT start_time, end_time FROM booking WHERE booking_id=?;"
	row := DB.QueryRow(query, bookingId)
	booking := &model.Booking{}
	row.Scan(&booking.StartTime, &booking.EndTime)
	return booking
}
