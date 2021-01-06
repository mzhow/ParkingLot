package dao

import (
	"ParkingLot/model"
	"database/sql"
	"errors"
	"time"
)

func GetBookingTime(bookingId int) *model.Booking {
	query := "SELECT start_time, end_time FROM booking WHERE booking_id=? AND valid=1"
	row := DB.QueryRow(query, bookingId)
	booking := &model.Booking{}
	row.Scan(&booking.StartTime, &booking.EndTime)
	return booking
}

func GetBookingCarAndSpot(bookingId int) (carId int, spotId int) {
	query := "SELECT car_id, spot_id FROM booking WHERE booking_id=? AND valid=1"
	row := DB.QueryRow(query, bookingId)
	row.Scan(&carId, &spotId)
	return carId, spotId
}

func UpdateBookingValid(bookingId int, valid int) error {
	update := "UPDATE booking SET valid=? WHERE booking_id=?"
	_, err := DB.Exec(update, valid, bookingId)
	return err
}

func GetBookingIdByCarAndSpot(carId int, spotId int) (bookingId int) {
	query := "SELECT booking_id FROM booking WHERE car_id=? AND spot_id=? AND valid=1"
	row := DB.QueryRow(query, carId, spotId)
	row.Scan(&bookingId)
	return bookingId
}

func InsertBooking(db *sql.DB, date string, carId int, spotId int, valid int) error {
	insert := "INSERT INTO booking(car_id,spot_id,start_time,end_time,valid)values(?,?,?,?,?)"
	layout := "2006-01-02 15:04:05"
	startTime, err := time.ParseInLocation(layout, date+" 08:00:00", time.Local)
	if err != nil {
		return err
	}
	endTime, err := time.ParseInLocation(layout, date+" 21:00:00", time.Local)
	if err != nil {
		return err
	}
	res, err := db.Exec(insert, carId, spotId, startTime, endTime, valid)
	if err != nil {
		return err
	}
	affect, _ := res.RowsAffected()
	if affect != 1 {
		return errors.New("dao: InsertBooking affect 0 row")
	}
	return nil
}
