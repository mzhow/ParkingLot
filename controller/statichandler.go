package controller

import (
	"ParkingLot/dao"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("Authorization")
	if checkToken(token) == false {
		// 如果token不正确则返回403页面
		forbiddenHandler(w)
		return
	}
	// 更新用户登录时间
	parseToken, err := ParseToken(token)
	checkErr(err)
	username := parseToken.Audience
	err = dao.UpdateLoginTime(username)
	checkErr(err)
	// 将用户详细信息发给用户
	user := dao.UserDetail(username)
	tmpl, err := template.ParseFiles("views/pages/user/index.html")
	checkErr(err)
	err = tmpl.Execute(w, user)
	checkErr(err)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		tmpl, err := template.ParseFiles("views/pages/user/login.html")
		checkErr(err)
		err = tmpl.Execute(w, nil)
		checkErr(err)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/pages/user/register.html")
	checkErr(err)
	err = tmpl.Execute(w, nil)
	checkErr(err)
}

func forbiddenHandler(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("views/pages/error/403.html")
	checkErr(err)
	err = tmpl.Execute(w, nil)
	checkErr(err)
}

func payHandler(w http.ResponseWriter, r *http.Request, username string) {
	tmpl, err := template.ParseFiles("views/pages/user/pay.html")
	checkErr(err)
	err = tmpl.Execute(w, dao.GetUserFee(username))
	checkErr(err)
}
