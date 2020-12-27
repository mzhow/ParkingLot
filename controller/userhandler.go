package controller

import (
	"ParkingLot/dao"
	"html/template"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// 进入index页面前先检查有没有session，有的话查里面包含的用户名和密码，直接进入已登录界面
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
		tmpl, err := template.ParseFiles("views/pages/user/index.html")
		checkErr(err)
		err = tmpl.Execute(w, nil)
		checkErr(err)
	}
}