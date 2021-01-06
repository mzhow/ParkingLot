package dao

import (
	"ParkingLot/model"
	"database/sql"
	"errors"
	"strconv"
	"time"
)

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

func GetTodaySpotCount(isEmpty int, indoor int) (count int) {
	query := "SELECT count(*) FROM spot WHERE is_empty=? AND indoor=? AND valid=1 AND spot_id NOT IN " +
		"(SELECT spot_id FROM booking WHERE valid=1 AND TO_DAYS(start_time)=TO_DAYS(NOW()))"
	row := DB.QueryRow(query, isEmpty, indoor)
	row.Scan(&count)
	return count
}

func GetTomorrowSpotCount(isEmpty int, indoor int) (count int) {
	query := "SELECT count(*) FROM spot WHERE is_empty=? AND indoor=? AND valid=1 AND spot_id NOT IN " +
		"(SELECT spot_id FROM booking WHERE valid=1 AND TO_DAYS(start_time)-TO_DAYS(NOW())=1)"
	row := DB.QueryRow(query, isEmpty, indoor)
	row.Scan(&count)
	return count
}

func GetRequiredSpot(db *sql.DB, date string, charging string, indoor string, outdoor string) (spotId int, err error) {
	var query string
	if indoor == "1" && outdoor == "1" {
		query = "SELECT spot_id FROM spot WHERE is_empty=1 AND valid=1 AND charging=? AND spot_id NOT IN " +
			"(SELECT spot_id FROM booking WHERE valid=1 AND TO_DAYS(start_time)=TO_DAYS(?)) ORDER BY rand() LIMIT 1"
	} else if indoor == "0" && outdoor == "1" {
		query = "SELECT spot_id FROM spot WHERE is_empty=1 AND valid=1 AND indoor=0 AND charging=? AND spot_id NOT IN " +
			"(SELECT spot_id FROM booking WHERE valid=1 AND TO_DAYS(start_time)=TO_DAYS(?)) ORDER BY rand() LIMIT 1"
	} else if indoor == "1" && outdoor == "0" {
		query = "SELECT spot_id FROM spot WHERE is_empty=1 AND valid=1 AND indoor=1 AND charging=? AND spot_id NOT IN " +
			"(SELECT spot_id FROM booking WHERE valid=1 AND TO_DAYS(start_time)=TO_DAYS(?)) ORDER BY rand() LIMIT 1"
	} else {
		return 0, errors.New("dao: GetRequiredSpot cannot run while indoor and outdoor both equal to 0")
	}
	dateTime, err := time.ParseInLocation("2006-01-02 15:04:05", date+" 08:00:00", time.Local)
	if err != nil {
		return 0, err
	}
	chargingInt, err := strconv.Atoi(charging)
	if err != nil {
		return 0, err
	}
	row := db.QueryRow(query, chargingInt, dateTime)
	err = row.Scan(&spotId)
	if err != nil {
		return 0, err
	}
	return spotId, nil
}
