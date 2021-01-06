package controller

import (
	"ParkingLot/dao"
	"html/template"
	"net/http"
	"strconv"
)

// 判断请求所带的用户名密码是否拥有roleId的权限
func checkAdminPermission(w http.ResponseWriter, r *http.Request, roleId int) bool {
	sess := globalSessions.SessionStart(w, r)
	sessUsername := sess.Get("username")
	sessPassword := sess.Get("password")
	if sessUsername != nil && sessPassword != nil {
		encodePassword, err := dao.GetAdminEncodePassword(sessUsername.(string))
		checkErr(err)
		if encodePassword != sessPassword.(string) {
			return false
		}
		if dao.CheckAdminRole(sessUsername.(string), roleId) == false {
			return false
		}
		return true
	}
	return false
}

// 处理管理员用户登录行为
func doAdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	captchaId := r.FormValue("captchaId")
	validateCode := r.FormValue("validateCode")
	if VerifyCaptcha(captchaId, validateCode) == false {
		logInfo(httpReqInfo(r) + " admin username: " + username + " input wrong validateCode")
		adminForbiddenHandler(w)
		return
	}
	encodePassword, err := dao.GetAdminEncodePassword(username)
	checkErr(err)
	if ComparePasswords(encodePassword, password) == false {
		logInfo(httpReqInfo(r) + " admin username: " + username + " input wrong password")
		adminForbiddenHandler(w)
		return
	}
	// 若当前请求cookie中存在有效的session，则返回那个session，否则就创建一个新session
	sess := globalSessions.SessionStart(w, r)
	val := sess.Get("username")
	if val != nil {
		logInfo(httpReqInfo(r) + " admin username: " + username + " session for user already exists")
	} else {
		// 将用户名和加密后的密码存入session
		sess.Set("username", username)
		sess.Set("password", encodePassword)
		logInfo(httpReqInfo(r) + " admin username: " + username + " set session for user")
	}
	err = dao.UpdateAdminLoginTime(username)
	checkErr(err)
	page := r.URL.Query().Get("page") //获取带有参数的给请求的url
	if page == "" {
		page = "1"
	}
	users, err := dao.GetUsers(page)
	checkErr(err)
	pagenum, err := strconv.Atoi(page)
	checkErr(err)
	pager := dao.CreatePaginator(pagenum, dao.GetDefaultPageSize(), dao.CountUsers())
	tmpl, err := template.ParseFiles("views/pages/admin/admin_index.html")
	checkErr(err)
	err = tmpl.Execute(w, map[string]interface{}{"users": users, "paginator": pager})
	checkErr(err)
}

// 管理员权限不足403页面
func adminForbiddenHandler(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("views/pages/admin/admin_403.html")
	checkErr(err)
	err = tmpl.Execute(w, nil)
	checkErr(err)
	return
}

// 管理员登录页面
func adminLoginHandler(w http.ResponseWriter, r *http.Request) {
	if checkAdminPermission(w, r, 2) {
		adminUsersDetailHandler(w, r)
		return
	}
	tmpl, err := template.ParseFiles("views/pages/admin/admin_login.html")
	checkErr(err)
	err = tmpl.Execute(w, nil)
	checkErr(err)
}

// 管理员账号登出
func adminLogoutHandler(w http.ResponseWriter, r *http.Request) {
	// 登出时删除session，回到登录界面
	sess := globalSessions.SessionStart(w, r)
	val := sess.Get("username")
	if val != nil {
		globalSessions.SessionDestroy(w, r)
		err := dao.UpdateAdminLogoutTime(val.(string))
		logInfo(httpReqInfo(r) + " username: " + val.(string) + " user logout")
		checkErr(err)
	}
	tmpl, err := template.ParseFiles("views/pages/admin/admin_login.html")
	checkErr(err)
	err = tmpl.Execute(w, nil)
	checkErr(err)
}

// 查看用户详情
func adminUsersDetailHandler(w http.ResponseWriter, r *http.Request) {
	if checkAdminPermission(w, r, 2) == false {
		adminForbiddenHandler(w)
		return
	}
	page := r.URL.Query().Get("page") //获取带有参数的给请求的url
	if page == "" {
		page = "1"
	}
	users, err := dao.GetUsers(page)
	checkErr(err)
	pagenum, err := strconv.Atoi(page)
	checkErr(err)
	pager := dao.CreatePaginator(pagenum, dao.GetDefaultPageSize(), dao.CountUsers())
	tmpl, err := template.ParseFiles("views/pages/admin/admin_index.html")
	checkErr(err)
	err = tmpl.Execute(w, map[string]interface{}{"users": users, "paginator": pager})
	checkErr(err)
}

// 查看车位详情
func adminSpotsDetailHandler(w http.ResponseWriter, r *http.Request) {
	if checkAdminPermission(w, r, 2) == false {
		adminForbiddenHandler(w)
		return
	}
	page := r.URL.Query().Get("page") //获取带有参数的给请求的url
	if page == "" {
		page = "1"
	}
	spots, err := dao.GetSpots(page)
	checkErr(err)
	pagenum, err := strconv.Atoi(page)
	checkErr(err)
	pager := dao.CreatePaginator(pagenum, dao.GetDefaultPageSize(), dao.CountSpots())
	tmpl, err := template.ParseFiles("views/pages/admin/admin_spots.html")
	checkErr(err)
	err = tmpl.Execute(w, map[string]interface{}{"spots": spots, "paginator": pager})
	checkErr(err)
}

// 查看订单详情
func adminBookingsDetailHandler(w http.ResponseWriter, r *http.Request) {
	if checkAdminPermission(w, r, 2) == false {
		adminForbiddenHandler(w)
		return
	}
	page := r.URL.Query().Get("page") //获取带有参数的给请求的url
	if page == "" {
		page = "1"
	}
	bookings, err := dao.GetBookings(page)
	checkErr(err)
	pagenum, err := strconv.Atoi(page)
	checkErr(err)
	pager := dao.CreatePaginator(pagenum, dao.GetDefaultPageSize(), dao.CountBookings())
	tmpl, err := template.ParseFiles("views/pages/admin/admin_bookings.html")
	checkErr(err)
	err = tmpl.Execute(w, map[string]interface{}{"bookings": bookings, "paginator": pager})
	checkErr(err)
}
