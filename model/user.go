package model

import "time"

type User struct {
	UserId     int
	RoleId     int
	CarId      int
	BookingId  int
	Username   string
	LoginTime  time.Time
	LogoutTime time.Time
	Fee        float32
	Valid      int
}

type UserDetail struct {
	UserId   int
	RoleName string
	Username string
	Fee      float32
	Car      *Car
	Spot     *Spot
	Booking  *Booking
}

/*
CREATE TABLE IF NOT EXISTS user (
	user_id INT UNSIGNED AUTO_INCREMENT COMMENT '用户ID',
	role_id INT NOT NULL COMMENT '角色ID(1-车主 2-管理员 3-超级管理员)',
	car_id INT NOT NULL COMMENT '用户车辆',
	booking_id INT COMMENT '预定订单号',
	username VARCHAR(50) NOT NULL COMMENT '用户名',
	password VARCHAR(100) NOT NULL COMMENT '加密后的密码',
	login_time DATETIME COMMENT '上次登录时间',
	logout_time DATETIME COMMENT '上次登出时间',
	fee FLOAT NOT NULL COMMENT '待支付费用',
	valid INT NOT NULL COMMENT '用户是否可下单',
    PRIMARY KEY (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS role (
	role_id INT UNSIGNED COMMENT '角色ID',
	role_name VARCHAR(50) NOT NULL COMMENT '角色名(1-车主 2-管理员 3-超级管理员)',
    PRIMARY KEY (role_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT role SET role_id=1, role_name='车主';
INSERT role SET role_id=2, role_name='管理员';
INSERT role SET role_id=3, role_name='超级管理员';
*/
