package controller

import (
	"ParkingLot/dao"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

// 检查用户名和密码如果正确则发给用户一个token
func doLoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	encodePassword, _ := dao.GetEncodePassword(username)
	res := ResData{}
	if ComparePasswords(encodePassword, password) == false {
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
	err = dao.InsertUser(1, dao.GetCarId(carName), username, HashAndSalt(password), 1)
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
