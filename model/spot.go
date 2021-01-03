package model

//type Spot struct {
//	SpotId      int
//	SpotName    string
//	SpotSection string
//	IsEmpty     int
//	Indoor      int
//	Charging    int
//	HourlyFee   int
//	Valid       int
//}

type Spot struct {
	SpotName string
	IsEmpty  int
	Indoor   int
	Charging int
	DailyFee float32
}

/*
CREATE TABLE IF NOT EXISTS spot (
	spot_id INT UNSIGNED AUTO_INCREMENT COMMENT '车位ID',
	spot_name VARCHAR(50) NOT NULL COMMENT '车位名',
	is_empty INT NOT NULL COMMENT '车位是否为空',
	indoor INT NOT NULL COMMENT '车位是否在室内',
	charging INT NOT NULL COMMENT '车位是否设置充电桩',
	daily_fee FLOAT NOT NULL COMMENT '停车每天价格',
	valid INT NOT NULL COMMENT '车位是否有效',
    PRIMARY KEY (spot_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
