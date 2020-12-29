package controller

import (
	"ParkingLot/dao"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// 进入login页面前先检查有没有session，有的话查里面包含的用户名和密码，直接进入已登录界面
	if r.URL.Path == "/" {
		sess := globalSessions.SessionStart(w, r)
		if sess.Get("username") != nil && sess.Get("password") != nil {
			loginUsername := sess.Get("username")
			loginPassword := sess.Get("password")
			encodePassword := dao.GetEncodePassword(loginUsername.(string))
			if ComparePasswords(encodePassword, loginPassword.(string)) {
				tmpl, err := template.ParseFiles("static/logined.html")
				checkErr(err)
				err = tmpl.Execute(w, map[string]interface{}{"user": dao.UserDetail(loginUsername.(string))})
				checkErr(err)
			}
			//if len(objs) > 0 {
			//	tmpl, err := template.ParseFiles("static/logined.html")
			//	checkErr(err)
			//	str := strconv.Itoa(objs[0].Userid)
			//	sql.UpdateLogintime("UPDATE user_logintime SET logintime=? WHERE userid=?", objs[0].Userid, sql.Timenow())
			//	err = tmpl.Execute(w, map[string]interface{}{"suc_msg": "您已通过session登录", "userinfo": objs[0], "userid": str})
			//	checkErr(err)
			//	//log.Print("indexHandler...  " + sess.Get("username").(string) + " logined by session successfully")
			//	return
			//}
		}
		tmpl, err := template.ParseFiles("views/pages/user/login.html")
		checkErr(err)
		err = tmpl.Execute(w, nil)
		checkErr(err)
	}
}

// 校验请求头中的Token，若正确返回对应的用户成功登录页面
func checkTokenHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "null" {
		return
	}
	parseToken, err := ParseToken(token)
	checkErr(err)
	// token错误
	if !dao.CheckUsernameValid(parseToken.Audience) {
		return
	}
	user := dao.UserDetail(parseToken.Audience)
	tmpl, err := template.ParseFiles("views/pages/user/index.html")
	checkErr(err)
	err = tmpl.Execute(w, user)
	checkErr(err)
}

func doLoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	encodePassword := dao.GetEncodePassword(username)
	res := JwtRes{}
	if ComparePasswords(encodePassword, password) == false {
		// 登录失败
		res = JwtRes{
			Valid: 0,
		}
	} else {
		// 登录成功
		res = JwtRes{
			Valid: 1,
			Token: CreateToken(username),
		}
	}
	jsonData, err := json.Marshal(res)
	checkErr(err)
	_, err = io.WriteString(w, string(jsonData))
	checkErr(err)
}

func forgotHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/pages/user/forgot.html")
	checkErr(err)
	err = tmpl.Execute(w, nil)
	checkErr(err)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/pages/user/register.html")
	checkErr(err)
	err = tmpl.Execute(w, nil)
	checkErr(err)
}

//doRegister
func doRegisterHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	carName := r.FormValue("car_name")
	log.Print("doRegisterHandler...  username: " + username + "  password: " + password)
	// username和password均不能为空
	if username == "" || len(password) < 8 {
		return
	}
	// 存在相同的用户名或车牌号
	if dao.ExistUsernameOrCarName(username, carName) {
		//tmpl, err := template.ParseFiles("static/registered.html")
		//checkErr(err)
		//err = tmpl.Execute(w, map[string] string {"err_msg": "已存在相同用户名或密码"})
		//checkErr(err)
		//log.Print("doRegisterHandler...  register failed, the username already exists")
		//return
	}

	// 向车辆表中添加车辆信息
	dao.InsertCar(carName, 0)
	// 向用户表中添加用户信息
	dao.InsertUser(1, dao.GetCarId(carName), username, password, 1)

	// 向登录时间表插入登录时间
	sql.InsertLoginTime("INSERT user_logintime SET userid=?, logintime=?", userid)
	tmpl, err := template.ParseFiles("static/registered.html")
	checkErr(err)
	objs := sql.QueryUser(username, password)
	str := strconv.Itoa(objs[0].Userid)
	err = tmpl.Execute(w, map[string]interface{}{"suc_msg": "欢迎注册", "userinfo": objs[0], "userid": str})
	checkErr(err)
	log.Print("doRegisterHandler...  " + username + " registered successfully")
}
