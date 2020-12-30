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
	SpotName  string
	IsEmpty   int
	Indoor    int
	Charging  int
	HourlyFee int
}

/*
CREATE TABLE IF NOT EXISTS spot (
	spot_id INT UNSIGNED AUTO_INCREMENT COMMENT '车位ID',
	spot_name VARCHAR(50) NOT NULL COMMENT '车位名',
	spot_section VARCHAR(50) NOT NULL COMMENT '车位所在区域',
	is_empty INT NOT NULL COMMENT '车位是否为空',
	indoor INT NOT NULL COMMENT '车位是否在室内',
	charging INT NOT NULL COMMENT '车位是否设置充电桩',
	hourly_fee INT NOT NULL COMMENT '停车每小时单价',
	valid INT NOT NULL COMMENT '车位是否有效',
    PRIMARY KEY (spot_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
