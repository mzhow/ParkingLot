package controller

import (
	"ParkingLot/dao"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("Authorization")
	if checkToken(token) == false {
		tmpl, err := template.ParseFiles("views/pages/error/404.html")
		checkErr(err)
		err = tmpl.Execute(w, nil)
		checkErr(err)
		return
	}
	// 更新用户登录时间
	parseToken, err := ParseToken(token)
	checkErr(err)
	username := parseToken.Audience
	err = dao.UpdateLoginTime(username)
	checkErr(err)
	// 将用户详细信息发给用户
	w.Header().Set("content-type", "text/html")
	user := dao.UserDetail(username)
	tmpl, err := template.ParseFiles("views/pages/user/index.html")
	checkErr(err)
	err = tmpl.Execute(w, user)
	checkErr(err)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/pages/user/login.html")
	checkErr(err)
	err = tmpl.Execute(w, nil)
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
