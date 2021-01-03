package model

import "time"

//type Car struct {
//	CarId     int
//	CarNumber string
//	IsParking int
//	EntryTime time.Time
//	OutTime   time.Time
//}

type Car struct {
	CarName   string
	IsParking int
	EntryTime time.Time
	OutTime   time.Time
}

/*
CREATE TABLE IF NOT EXISTS car (
	car_id INT UNSIGNED AUTO_INCREMENT COMMENT '车辆ID',
	car_name VARCHAR(50) NOT NULL COMMENT '车牌号',
	is_parking INT NOT NULL COMMENT '是否已入库',
	entry_time DATETIME COMMENT '入库时间',
	out_time   DATETIME COMMENT '出库时间',
    PRIMARY KEY (car_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
