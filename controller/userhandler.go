package controller

import (
	"ParkingLot/dao"
	"ParkingLot/model"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func httpReqInfo(r *http.Request) string {
	return "Request IP: " + r.RemoteAddr
}

// 校验请求头中的Token，若正确返回valid=1否则返回valid=0
func checkTokenHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	res := ResData{}
	if checkToken(token) {
		res = ResData{
			Valid: 1,
		}
	} else {
		res = ResData{
			Valid: 0,
		}
		logInfo(httpReqInfo(r) + " check token failed")
	}
	// 返回客户端验证结果
	jsonData, err := json.Marshal(res)
	checkErr(err)
	_, err = io.WriteString(w, string(jsonData))
	checkErr(err)
}

// 获取验证码
func getCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	id, b64s := GetCaptcha()
	res := Captcha{
		Id:   id,
		B64s: b64s,
	}
	// 返回客户端验证结果
	jsonData, err := json.Marshal(res)
	checkErr(err)
	_, err = io.WriteString(w, string(jsonData))
	checkErr(err)
}

// 获取停车场数量信息
func getSpotHandler(w http.ResponseWriter, r *http.Request) {
	res := model.SpotCount{
		TodayIndoor:     dao.GetTodaySpotCount(1, 1),
		TodayOutdoor:    dao.GetTodaySpotCount(1, 0),
		TomorrowIndoor:  dao.GetTomorrowSpotCount(1, 1),
		TomorrowOutdoor: dao.GetTomorrowSpotCount(1, 0),
	}
	// 返回客户端验证结果
	jsonData, err := json.Marshal(res)
	checkErr(err)
	_, err = io.WriteString(w, string(jsonData))
	checkErr(err)
}

// 检查用户名和密码如果正确则发给用户一个token
func doLoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	captchaId := r.FormValue("captchaId")
	validateCode := r.FormValue("validateCode")
	encodePassword, _ := dao.GetEncodePassword(username)
	res := ResData{}
	if VerifyCaptcha(captchaId, validateCode) == false {
		// 验证码错误
		logInfo(httpReqInfo(r) + " username: " + username + " input wrong validateCode")
		res = ResData{
			Valid:   0,
			Message: "验证码错误",
		}
	} else if ComparePasswords(encodePassword, password) == false {
		// 登录失败
		logInfo(httpReqInfo(r) + " username: " + username + " input wrong password")
		res = ResData{
			Valid:   0,
			Message: "用户名或密码错误",
		}
	} else {
		// 登录成功
		logInfo(httpReqInfo(r) + " username: " + username + " verify user success")
		res = ResData{
			Valid: 1,
			Token: CreateToken(username),
		}
	}
	// 返回客户端验证结果
	jsonData, err := json.Marshal(res)
	checkErr(err)
	_, err = io.WriteString(w, string(jsonData))
	checkErr(err)
}

// 用户注册
func doRegisterHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	carName := r.FormValue("car_name")
	logInfo(httpReqInfo(r) + " username: " + username + " do register")
	// username和password均不能为空，密码长度应不少于8
	if username == "" || len(password) < 8 {
		return
	}
	// 存在相同的用户名或车牌号
	if dao.ExistUsernameOrCarName(username, carName) {
		res := ResData{
			Valid:   0,
			Message: "已存在相同的用户名或车牌号",
		}
		logInfo(httpReqInfo(r) + " username: " + username + " register failed, the same username or car_name already exists")
		// 返回客户端注册失败结果
		jsonData, err := json.Marshal(res)
		checkErr(err)
		_, err = io.WriteString(w, string(jsonData))
		checkErr(err)
		return
	}
	// 向车辆表中添加车辆信息
	err := dao.InsertCar(carName, 0)
	logInfo(httpReqInfo(r) + " username: " + username + " register，add car: " + carName)
	checkErr(err)
	// 向用户表中添加用户信息
	err = dao.InsertUser(1, dao.GetCarIdByCarName(carName), username, HashAndSalt(password), 0, 1)
	logInfo(httpReqInfo(r) + " username: " + username + " add user info")
	checkErr(err)
	// 更新登录登出时间
	err = dao.UpdateLoginTime(username)
	checkErr(err)
	err = dao.UpdateLogoutTime(username)
	checkErr(err)
	res := ResData{
		Valid: 1,
		Token: CreateToken(username),
	}
	// 返回客户端注册成功结果
	jsonData, err := json.Marshal(res)
	checkErr(err)
	_, err = io.WriteString(w, string(jsonData))
	checkErr(err)
}

// 用户登出
func doLogoutHandler(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("Authorization")
	if checkToken(token) == false {
		logInfo(httpReqInfo(r) + " check token failed")
		// 如果token不正确则返回403页面
		forbiddenHandler(w)
		return
	}
	// 更新用户登出时间
	parseToken, err := ParseToken(token)
	checkErr(err)
	username := parseToken.Audience
	err = dao.UpdateLogoutTime(username)
	logInfo(httpReqInfo(r) + " username: " + username + " user logout")
	checkErr(err)
	// 返回主页
	r.URL.Path = "/"
	loginHandler(w, r)
}

// 进入停车场
func entryHandler(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("Authorization")
	if checkToken(token) == false {
		logInfo(httpReqInfo(r) + " check token failed")
		// 如果token不正确则返回403页面
		forbiddenHandler(w)
		return
	}
	parseToken, err := ParseToken(token)
	checkErr(err)
	username := parseToken.Audience
	bookingId := dao.GetBookingId(username)
	carId, spotId := dao.GetBookingCarAndSpot(bookingId)
	// 判断是否已进入车位或者不在订单预定时间范围内
	if dao.GetCar(carId).IsParking == 1 ||
		time.Now().After(dao.GetBookingTime(bookingId).EndTime) ||
		time.Now().Before(dao.GetBookingTime(bookingId).StartTime) {
		indexHandler(w, r)
		return
	}
	// 如果有未缴费用，跳转支付页面
	if dao.GetUserFee(username) != 0 {
		// 进入支付页面
		payHandler(w, r, username)
		return
	}
	err = dao.UpdateSpot(spotId, 0)
	logInfo(httpReqInfo(r) + " username: " + username + " update spot")
	checkErr(err)
	err = dao.UpdateCarEntryTime(carId)
	logInfo(httpReqInfo(r) + " username: " + username + " update car entry_time")
	checkErr(err)
	// 刷新页面
	indexHandler(w, r)
}

// 离开停车场
func outHandler(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("Authorization")
	if checkToken(token) == false {
		logInfo(httpReqInfo(r) + " check token failed")
		// 如果token不正确则返回403页面
		forbiddenHandler(w)
		return
	}
	parseToken, err := ParseToken(token)
	checkErr(err)
	username := parseToken.Audience
	bookingId := dao.GetBookingId(username)
	carId, spotId := dao.GetBookingCarAndSpot(bookingId)
	// 判断是否未进入车位或当前还未到订单开始时间
	if dao.GetCar(carId).IsParking == 0 ||
		time.Now().Before(dao.GetBookingTime(bookingId).StartTime) {
		indexHandler(w, r)
		return
	}
	// 如果超时，重新更新费用
	if time.Now().After(dao.GetBookingTime(bookingId).EndTime.Add(time.Hour)) {
		extraHours := time.Since(dao.GetBookingTime(bookingId).EndTime.Add(time.Hour)).Hours()
		extraDays := int(extraHours/24 + 1) // 超过22:00算一整天
		fee := dao.GetSpotDailyFee(spotId) * float32(extraDays)
		err = dao.UpdateUserFee(username, fee) // 更新用户产生的超时费用
		logInfo(httpReqInfo(r) + " username: " + username + " update extra fee")
		checkErr(err)
		// 进入支付页面
		payHandler(w, r, username)
		return
	}
	// 不超时直接出停车场
	err = dao.UpdateSpot(spotId, 1)
	logInfo(httpReqInfo(r) + " username: " + username + " update spot")
	checkErr(err)
	err = dao.UpdateCarOutTime(carId)
	logInfo(httpReqInfo(r) + " username: " + username + " update car out_time")
	checkErr(err)
	// 刷新页面
	indexHandler(w, r)
}

// 订单请求处理
func bookingHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	captchaId := r.FormValue("captchaId")
	validateCode := r.FormValue("validateCode")
	bookingDate := r.FormValue("bookingDate")
	charging := r.FormValue("needCharging")
	indoor := r.FormValue("chooseIndoor")
	outdoor := r.FormValue("chooseOutdoor")
	res := ResData{}
	if checkToken(token) == false {
		logInfo(httpReqInfo(r) + " check token failed")
		res = ResData{
			Valid:   0,
			Message: "用户登录身份过期",
		}
		// 返回客户端验证结果
		jsonData, err := json.Marshal(res)
		checkErr(err)
		_, err = io.WriteString(w, string(jsonData))
		checkErr(err)
		return
	}
	parseToken, err := ParseToken(token)
	checkErr(err)
	username := parseToken.Audience
	if dao.CheckCarIsParking(dao.GetCarIdByUsername(username)) == true {
		res = ResData{
			Valid:   0,
			Message: "已有车辆在停车场内，不允许新建订单",
		}
	} else if dao.GetUserFee(username) != 0 {
		res = ResData{
			Valid:   0,
			Message: "还有待支付的费用",
		}
	} else if VerifyCaptcha(captchaId, validateCode) == false {
		res = ResData{
			Valid:   0,
			Message: "验证码错误",
		}
	} else if bookingDate == time.Now().Format("2006-01-02") && time.Now().Hour() >= 21 {
		res = ResData{
			Valid:   0,
			Message: "已不能预约今日车位",
		}
	} else if bookingDate == time.Now().Add(time.Hour*24).Format("2006-01-02") && time.Now().Hour() <= 21 {
		res = ResData{
			Valid:   0,
			Message: "22:00开放明日车位预约",
		}
	} else if indoor == "0" && outdoor == "0" {
		res = ResData{
			Valid:   0,
			Message: "室内室外至少选一种",
		}
	} else {
		// 满足下单要求
		res = ResData{
			Valid:   1,
			Message: "正在处理您的订单",
		}
		req := BookingRequest{
			Username: username,
			Date:     bookingDate,
			Charging: charging,
			Indoor:   indoor,
			Outdoor:  outdoor,
		}
		rs, err := json.Marshal(req)
		checkErr(err)
		logInfo(httpReqInfo(r) + " username: " + username + " start to deal with booking")
		RabbitMQSend(rs) // 将订单请求发送到RabbitMQ
	}
	// 返回用户验证结果
	jsonData, err := json.Marshal(res)
	checkErr(err)
	_, err = io.WriteString(w, string(jsonData))
	checkErr(err)
}

// 处理订单请求
func makeBooking(req BookingRequest) error {
	tx, err := dao.DB.Begin()
	if err != nil {
		return err
	}
	err = dao.UpdateBookingValid(dao.GetBookingId(req.Username), 0)
	logInfo("makeBooking... username: " + req.Username + " update booking valid to 0")
	if err != nil {
		return err
	}
	err = dao.UpdateUserValid(req.Username, 1)
	logInfo("makeBooking... username: " + req.Username + " update user valid to 1")
	if err != nil {
		return err
	}
	spotId, err := dao.GetRequiredSpot(dao.DB, req.Date, req.Charging, req.Indoor, req.Outdoor)
	if err != nil {
		logInfo("makeBooking... username: " + req.Username + " cannot get required spot")
		if err := tx.Rollback(); err != nil {
			checkErr(err)
		}
		return err
	}
	carId := dao.GetCarIdByUsername(req.Username)
	err = dao.InsertBooking(dao.DB, req.Date, carId, spotId, 1)
	logInfo("makeBooking... username: " + req.Username + " insert booking")
	if err != nil {
		logInfo("makeBooking... username: " + req.Username + " cannot insert booking")
		if err := tx.Rollback(); err != nil {
			checkErr(err)
		}
		return err
	}
	bookingId := dao.GetBookingIdByCarAndSpot(carId, spotId)
	err = dao.UpdateUserForNewBooking(req.Username, bookingId, dao.GetSpotDailyFee(spotId))
	logInfo("makeBooking... username: " + req.Username + " update user for new booking")
	if err != nil {
		logInfo("makeBooking... username: " + req.Username + " cannot update user for new booking")
		if err := tx.Rollback(); err != nil {
			checkErr(err)
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func cancelBookingHandler(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("Authorization")
	if checkToken(token) == false {
		logInfo(httpReqInfo(r) + " check token failed")
		// 如果token不正确则返回403页面
		forbiddenHandler(w)
		return
	}
	parseToken, err := ParseToken(token)
	checkErr(err)
	username := parseToken.Audience
	bookingId := dao.GetBookingId(username)
	// 仅允许在订单开始时间前取消订单
	if time.Now().Before(dao.GetBookingTime(bookingId).StartTime) {
		logInfo(httpReqInfo(r) + " username: " + username + " cancel booking")
		err := dao.UpdateBookingValid(bookingId, 0)
		checkErr(err)
		err = dao.UpdateUserFee(username, 0)
		checkErr(err)
		err = dao.UpdateUserValid(username, 1)
		checkErr(err)
		indexHandler(w, r)
		return
	}
	forbiddenHandler(w)
}

func payFeeHandler(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("Authorization")
	if checkToken(token) == false {
		logInfo(httpReqInfo(r) + " check token failed")
		// 如果token不正确则返回403页面
		forbiddenHandler(w)
		return
	}
	parseToken, err := ParseToken(token)
	checkErr(err)
	username := parseToken.Audience
	err = dao.UpdateUserFee(username, 0)
	logInfo(httpReqInfo(r) + " username: " + username + " pay fee")
	checkErr(err)
	if r.FormValue("PathName") == "/out" {
		bookingId := dao.GetBookingId(username)
		carId, spotId := dao.GetBookingCarAndSpot(bookingId)
		err = dao.UpdateSpot(spotId, 1)
		logInfo(httpReqInfo(r) + " username: " + username + " update spot")
		checkErr(err)
		err = dao.UpdateCarOutTime(carId)
		logInfo(httpReqInfo(r) + " username: " + username + " update car out_time")
		checkErr(err)
		err = dao.UpdateBookingValid(dao.GetBookingId(username), 0)
		logInfo(httpReqInfo(r) + " username: " + username + " update booking valid to 0")
	}
	indexHandler(w, r)
}
