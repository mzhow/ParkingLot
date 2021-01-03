package model

import "time"

type AdminDetail struct {
	UserId     int
	RoleId     int
	Username   string
	LoginTime  time.Time
	LogoutTime time.Time
}

/*
CREATE TABLE IF NOT EXISTS admin (
	user_id INT UNSIGNED AUTO_INCREMENT COMMENT '用户ID',
	role_id INT NOT NULL COMMENT '角色ID(1-车主 2-管理员 3-超级管理员)',
	username VARCHAR(50) NOT NULL COMMENT '用户名',
	password VARCHAR(100) NOT NULL COMMENT '加密后的密码',
	login_time DATETIME COMMENT '上次登录时间',
	logout_time DATETIME COMMENT '上次登出时间',
	valid INT NOT NULL COMMENT '用户是否有效',
    PRIMARY KEY (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
