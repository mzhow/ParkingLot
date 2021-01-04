package controller

import (
	"ParkingLot/dao"
	"ParkingLot/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

//func loginHandler(w http.ResponseWriter, r *http.Request) {
//	// 进入login页面前先检查有没有session，有的话查里面包含的用户名和密码，直接进入已登录界面
//	if r.URL.Path == "/" {
//		sess := globalSessions.SessionStart(w, r)
//		if sess.Get("username") != nil && sess.Get("password") != nil {
//			loginUsername := sess.Get("username")
//			loginPassword := sess.Get("password")
//			encodePassword, _ := dao.GetEncodePassword(loginUsername.(string))
//			if ComparePasswords(encodePassword, loginPassword.(string)) {
//				tmpl, err := template.ParseFiles("static/logined.html")
//				checkErr(err)
//				err = tmpl.Execute(w, map[string]interface{}{"user": dao.UserDetail(loginUsername.(string))})
//				checkErr(err)
//			}
//			//if len(objs) > 0 {
//			//	tmpl, err := template.ParseFiles("static/logined.html")
//			//	checkErr(err)
//			//	str := strconv.Itoa(objs[0].Userid)
//			//	sql.UpdateLogintime("UPDATE user_logintime SET logintime=? WHERE userid=?", objs[0].Userid, sql.Timenow())
//			//	err = tmpl.Execute(w, map[string]interface{}{"suc_msg": "您已通过session登录", "userinfo": objs[0], "userid": str})
//			//	checkErr(err)
//			//	//log.Print("indexHandler...  " + sess.Get("username").(string) + " logined by session successfully")
//			//	return
//			//}
//		}
//		tmpl, err := template.ParseFiles("views/pages/user/login.html")
//		checkErr(err)
//		err = tmpl.Execute(w, nil)
//		checkErr(err)
//	}
//}

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
	}
	// 返回客户端验证结果
	jsonData, err := json.Marshal(res)
	checkErr(err)
	_, err = io.WriteString(w, string(jsonData))
	checkErr(err)
}

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
		res = ResData{
			Valid:   0,
			Message: "验证码错误",
		}
	} else if ComparePasswords(encodePassword, password) == false {
		// 登录失败
		res = ResData{
			Valid:   0,
			Message: "用户名或密码错误",
		}
	} else {
		// 登录成功
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

//doRegister
func doRegisterHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	carName := r.FormValue("car_name")
	log.Print("doRegisterHandler...  username: " + username + "  password: " + password)
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
		// 返回客户端注册失败结果
		jsonData, err := json.Marshal(res)
		checkErr(err)
		_, err = io.WriteString(w, string(jsonData))
		checkErr(err)
		return
	}
	// 向车辆表中添加车辆信息
	err := dao.InsertCar(carName, 0)
	checkErr(err)
	// 向用户表中添加用户信息
	err = dao.InsertUser(1, dao.GetCarIdByCarName(carName), username, HashAndSalt(password), 0, 1)
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

func doLogoutHandler(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("Authorization")
	if checkToken(token) == false {
		// 如果token不正确则返回403页面
		forbiddenHandler(w)
		return
	}
	// 更新用户登出时间
	parseToken, err := ParseToken(token)
	checkErr(err)
	username := parseToken.Audience
	err = dao.UpdateLogoutTime(username)
	checkErr(err)
	// 返回主页
	loginHandler(w, r)
}

func entryHandler(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("Authorization")
	if checkToken(token) == false {
		// 如果token不正确则返回403页面
		forbiddenHandler(w)
		return
	}
	parseToken, err := ParseToken(token)
	checkErr(err)
	username := parseToken.Audience
	bookingId := dao.GetBookingId(username)
	carId, spotId := dao.GetBookingCarAndSpot(bookingId)
	// 判断是否已进入车位或者是否错过了预定时间
	if dao.GetCar(carId).IsParking == 1 ||
		time.Now().After(dao.GetBookingTime(bookingId).EndTime) ||
		time.Now().Before(dao.GetBookingTime(bookingId).StartTime) {
		indexHandler(w, r)
		return
	}
	err = dao.UpdateBookingValid(bookingId, 1)
	checkErr(err)
	err = dao.UpdateSpot(spotId, 0)
	checkErr(err)
	err = dao.UpdateCarEntryTime(carId)
	checkErr(err)
	// 刷新页面
	indexHandler(w, r)
}

func outHandler(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("Authorization")
	if checkToken(token) == false {
		// 如果token不正确则返回403页面
		forbiddenHandler(w)
		return
	}
	parseToken, err := ParseToken(token)
	checkErr(err)
	username := parseToken.Audience
	bookingId := dao.GetBookingId(username)
	carId, spotId := dao.GetBookingCarAndSpot(bookingId)
	// 判断是否未进入车位
	if dao.GetCar(carId).IsParking == 0 ||
		time.Now().Before(dao.GetBookingTime(bookingId).StartTime) {
		indexHandler(w, r)
		return
	}
	// 如果超时重新计算金额
	if time.Now().After(dao.GetBookingTime(bookingId).EndTime.Add(time.Hour)) {
		extraHours := time.Since(dao.GetBookingTime(bookingId).EndTime.Add(time.Hour)).Hours()
		extraDays := int(extraHours/24 + 1) // 超过22:00算一整天
		fee := dao.GetSpotDailyFee(spotId) * (float32(extraDays) + 1)
		err = dao.UpdateUserFee(username, fee) // 更新用户产生的超时费用
		checkErr(err)
		// 刷新页面
		indexHandler(w, r)
		return
	}
	err = dao.UpdateBookingValid(bookingId, 0)
	checkErr(err)
	err = dao.UpdateSpot(spotId, 1)
	checkErr(err)
	err = dao.UpdateCarOutTime(carId)
	checkErr(err)
	// 刷新页面
	indexHandler(w, r)
}

func bookingHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	captchaId := r.FormValue("captchaId")
	validateCode := r.FormValue("validateCode")
	bookingDate := r.FormValue("bookingDate")
	res := ResData{}
	if checkToken(token) == false {
		res = ResData{
			Valid:   0,
			Message: "用户登录身份过期或有未完成订单",
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
	} else {
		res = ResData{
			Valid:   1,
			Message: "正在处理您的订单",
		}
		parseToken, err := ParseToken(token)
		checkErr(err)
		username := parseToken.Audience
		req := BookingRequest{
			Username: username,
			Date:     bookingDate,
			Charging: r.FormValue("needCharging"),
			Indoor:   r.FormValue("chooseIndoor"),
			Outdoor:  r.FormValue("chooseOutdoor"),
		}
		rs, err := json.Marshal(req)
		checkErr(err)
		RabbitMQSend(rs)
	}
	// 返回客户端验证结果
	jsonData, err := json.Marshal(res)
	checkErr(err)
	_, err = io.WriteString(w, string(jsonData))
	checkErr(err)
}

func makeBooking(req BookingRequest) error {
	tx, err := dao.DB.Begin()
	if err != nil {
		return err
	}
	spotId, err := dao.GetRequiredSpot(dao.DB, req.Date, req.Charging, req.Indoor, req.Outdoor)
	if err != nil {
		log.Print("makeBooking...  cannot get required spot")
		if err := tx.Rollback(); err != nil {
			checkErr(err)
		}
		return err
	}
	carId := dao.GetCarIdByUsername(req.Username)
	err = dao.InsertBooking(dao.DB, req.Date, carId, spotId)
	if err != nil {
		log.Print("makeBooking...  cannot insert booking")
		if err := tx.Rollback(); err != nil {
			checkErr(err)
		}
		return err
	}
	bookingId := dao.GetBookingIdByCarAndSpot(carId, spotId)
	err = dao.UpdateUserForNewBooking(req.Username, bookingId, dao.GetSpotDailyFee(spotId))
	if err != nil {
		log.Print("makeBooking...  cannot update user for new booking")
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
