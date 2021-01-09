
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin`  (
  `user_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `role_id` int NOT NULL COMMENT '角色ID(1-车主 2-管理员 3-超级管理员)',
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '加密后的密码',
  `login_time` datetime(0) NULL DEFAULT NULL COMMENT '上次登录时间',
  `logout_time` datetime(0) NULL DEFAULT NULL COMMENT '上次登出时间',
  `valid` int NOT NULL COMMENT '用户是否有效',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES (1, 2, 'admin01', '$2a$10$FuexrhxcHGv/dRkVOX6ZUejAgcW.NAMVSEAAg2woB/kazPdQqvLqm', '2021-01-08 15:51:52', '2021-01-08 15:51:10', 1);

-- ----------------------------
-- Table structure for booking
-- ----------------------------
DROP TABLE IF EXISTS `booking`;
CREATE TABLE `booking`  (
  `booking_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单ID',
  `car_id` int UNSIGNED NOT NULL COMMENT '车辆ID',
  `spot_id` int UNSIGNED NOT NULL COMMENT '车辆ID',
  `start_time` datetime(0) NOT NULL COMMENT '开始时间',
  `end_time` datetime(0) NOT NULL COMMENT '结束时间',
  `valid` int NOT NULL COMMENT '订单是否有效：下单时赋1，取消或完成后赋0',
  PRIMARY KEY (`booking_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of booking
-- ----------------------------
INSERT INTO `booking` VALUES (1, 1, 10, '2021-01-08 08:00:00', '2021-01-08 21:00:00', 0);
INSERT INTO `booking` VALUES (2, 1, 48, '2021-01-09 08:00:00', '2021-01-09 21:00:00', 0);
INSERT INTO `booking` VALUES (3, 1, 34, '2021-01-09 08:00:00', '2021-01-09 21:00:00', 1);

-- ----------------------------
-- Table structure for car
-- ----------------------------
DROP TABLE IF EXISTS `car`;
CREATE TABLE `car`  (
  `car_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '车辆ID',
  `car_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '车牌号',
  `is_parking` int NOT NULL COMMENT '是否已入库',
  `entry_time` datetime(0) NULL DEFAULT NULL COMMENT '入库时间',
  `out_time` datetime(0) NULL DEFAULT NULL COMMENT '出库时间',
  PRIMARY KEY (`car_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of car
-- ----------------------------
INSERT INTO `car` VALUES (1, '京A1B666', 0, '2021-01-09 11:23:13', '2021-01-09 11:23:14');

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `role_id` int UNSIGNED NOT NULL COMMENT '角色ID',
  `role_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色名(1-车主 2-管理员 3-超级管理员)',
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, '车主');
INSERT INTO `role` VALUES (2, '管理员');
INSERT INTO `role` VALUES (3, '超级管理员');

-- ----------------------------
-- Table structure for spot
-- ----------------------------
DROP TABLE IF EXISTS `spot`;
CREATE TABLE `spot`  (
  `spot_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '车位ID',
  `spot_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '车位名',
  `is_empty` int NOT NULL COMMENT '车位是否为空',
  `indoor` int NOT NULL COMMENT '车位是否在室内',
  `charging` int NOT NULL COMMENT '车位是否设置充电桩',
  `daily_fee` float NOT NULL COMMENT '停车每天价格',
  `valid` int NOT NULL COMMENT '车位是否有效',
  PRIMARY KEY (`spot_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 51 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of spot
-- ----------------------------
INSERT INTO `spot` VALUES (1, 'A-01', 1, 1, 1, 60, 1);
INSERT INTO `spot` VALUES (2, 'A-02', 1, 1, 1, 60, 1);
INSERT INTO `spot` VALUES (3, 'A-03', 1, 1, 0, 60, 1);
INSERT INTO `spot` VALUES (4, 'A-04', 1, 1, 0, 60, 1);
INSERT INTO `spot` VALUES (5, 'A-05', 1, 0, 1, 50, 1);
INSERT INTO `spot` VALUES (6, 'A-06', 1, 0, 1, 50, 1);
INSERT INTO `spot` VALUES (7, 'A-07', 1, 0, 1, 50, 1);
INSERT INTO `spot` VALUES (8, 'A-08', 1, 0, 0, 50, 1);
INSERT INTO `spot` VALUES (9, 'A-09', 1, 0, 0, 50, 1);
INSERT INTO `spot` VALUES (10, 'A-10', 1, 0, 0, 50, 1);
INSERT INTO `spot` VALUES (11, 'B-01', 1, 1, 1, 60, 1);
INSERT INTO `spot` VALUES (12, 'B-02', 1, 1, 1, 60, 1);
INSERT INTO `spot` VALUES (13, 'B-03', 1, 1, 0, 60, 1);
INSERT INTO `spot` VALUES (14, 'B-04', 1, 1, 0, 60, 1);
INSERT INTO `spot` VALUES (15, 'B-05', 1, 0, 0, 50, 1);
INSERT INTO `spot` VALUES (16, 'B-06', 1, 0, 1, 50, 1);
INSERT INTO `spot` VALUES (17, 'B-07', 1, 0, 1, 50, 1);
INSERT INTO `spot` VALUES (18, 'B-08', 1, 0, 1, 50, 1);
INSERT INTO `spot` VALUES (19, 'B-09', 1, 0, 0, 50, 1);
INSERT INTO `spot` VALUES (20, 'B-10', 1, 0, 0, 50, 1);
INSERT INTO `spot` VALUES (21, 'B-11', 1, 0, 1, 50, 1);
INSERT INTO `spot` VALUES (22, 'B-12', 1, 0, 0, 50, 1);
INSERT INTO `spot` VALUES (23, 'C-01', 1, 1, 1, 60, 1);
INSERT INTO `spot` VALUES (24, 'C-02', 1, 1, 1, 60, 1);
INSERT INTO `spot` VALUES (25, 'C-03', 1, 1, 1, 60, 1);
INSERT INTO `spot` VALUES (26, 'C-04', 1, 1, 0, 60, 1);
INSERT INTO `spot` VALUES (27, 'C-05', 1, 1, 1, 60, 1);
INSERT INTO `spot` VALUES (28, 'C-06', 1, 1, 1, 60, 1);
INSERT INTO `spot` VALUES (29, 'C-07', 1, 1, 0, 60, 1);
INSERT INTO `spot` VALUES (30, 'C-08', 1, 1, 1, 60, 1);
INSERT INTO `spot` VALUES (31, 'C-09', 1, 1, 0, 60, 1);
INSERT INTO `spot` VALUES (32, 'C-10', 1, 1, 0, 60, 1);
INSERT INTO `spot` VALUES (33, 'C-11', 1, 1, 1, 60, 1);
INSERT INTO `spot` VALUES (34, 'C-12', 1, 0, 0, 50, 1);
INSERT INTO `spot` VALUES (35, 'C-13', 1, 0, 1, 50, 1);
INSERT INTO `spot` VALUES (36, 'C-14', 1, 0, 0, 50, 1);
INSERT INTO `spot` VALUES (37, 'C-15', 1, 0, 0, 50, 1);
INSERT INTO `spot` VALUES (38, 'C-16', 1, 0, 1, 50, 1);
INSERT INTO `spot` VALUES (39, 'C-17', 1, 0, 1, 50, 1);
INSERT INTO `spot` VALUES (40, 'C-18', 1, 0, 1, 50, 1);
INSERT INTO `spot` VALUES (41, 'C-19', 1, 0, 0, 50, 1);
INSERT INTO `spot` VALUES (42, 'C-20', 1, 0, 1, 50, 1);
INSERT INTO `spot` VALUES (43, 'C-21', 1, 0, 0, 50, 1);
INSERT INTO `spot` VALUES (44, 'C-22', 1, 0, 0, 50, 1);
INSERT INTO `spot` VALUES (45, 'C-23', 1, 0, 1, 50, 1);
INSERT INTO `spot` VALUES (46, 'C-24', 1, 0, 0, 50, 1);
INSERT INTO `spot` VALUES (47, 'C-25', 1, 0, 1, 50, 1);
INSERT INTO `spot` VALUES (48, 'C-26', 1, 0, 0, 50, 1);
INSERT INTO `spot` VALUES (49, 'C-27', 1, 0, 1, 50, 1);
INSERT INTO `spot` VALUES (50, 'C-28', 1, 0, 1, 50, 1);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `user_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `role_id` int NOT NULL COMMENT '角色ID(1-车主 2-管理员 3-超级管理员)',
  `car_id` int NOT NULL COMMENT '用户车辆',
  `booking_id` int NULL DEFAULT NULL COMMENT '预定订单号',
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '加密后的密码',
  `login_time` datetime(0) NULL DEFAULT NULL COMMENT '上次登录时间',
  `logout_time` datetime(0) NULL DEFAULT NULL COMMENT '上次登出时间',
  `fee` float NOT NULL COMMENT '待支付费用',
  `valid` int NOT NULL COMMENT '用户是否可下单',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 1, 1, 3, 'mzh', '$2a$10$UCtplCyZ02VOZG.LTQAfGO12aQbyfZT5OboR8Ez8LU6kaNxlRdbVS', '2021-01-09 11:23:14', '2021-01-09 11:04:38', 0, 0);

SET FOREIGN_KEY_CHECKS = 1;
