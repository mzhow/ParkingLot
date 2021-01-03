package model

import "time"

//type Booking struct {
//	BookingId int
//	CarId     int
//	SpotId    int
//	StartTime time.Time
//	EndTime   time.Time
//}

type Booking struct {
	CarName   string
	SpotName  string
	StartTime time.Time
	EndTime   time.Time
}

/*
CREATE TABLE IF NOT EXISTS booking (
	booking_id INT UNSIGNED AUTO_INCREMENT COMMENT '订单ID',
	car_id INT UNSIGNED NOT NULL COMMENT '车辆ID',
	spot_id INT UNSIGNED NOT NULL COMMENT '车辆ID',
	start_time DATETIME NOT NULL COMMENT '开始时间',
	end_time DATETIME NOT NULL COMMENT '结束时间',
	valid INT NOT NULL COMMENT '订单是否有效',
    PRIMARY KEY (booking_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
