package controller

//func adminHandler(w http.ResponseWriter, r *http.Request) {
//	sess := globalSessions.SessionStart(w, r)
//	loginUsername := sess.Get("username")
//	loginPassword := sess.Get("password")
//	if loginUsername != nil && loginPassword != nil {
//		encodePassword := dao.GetEncodePassword(loginUsername.(string))
//		if ComparePasswords(encodePassword, loginPassword.(string)) {
//			tmpl, err := template.ParseFiles("static/admin.html")
//			checkErr(err)
//			err = tmpl.Execute(w, map[string]interface{}{"admin": dao.AdminDetail(loginUsername.(string))})
//			checkErr(err)
//		}
//		//if len(objs) > 0 {
//		//	tmpl, err := template.ParseFiles("static/logined.html")
//		//	checkErr(err)
//		//	str := strconv.Itoa(objs[0].Userid)
//		//	sql.UpdateLogintime("UPDATE user_logintime SET logintime=? WHERE userid=?", objs[0].Userid, sql.Timenow())
//		//	err = tmpl.Execute(w, map[string]interface{}{"suc_msg": "您已通过session登录", "userinfo": objs[0], "userid": str})
//		//	checkErr(err)
//		//	//log.Print("indexHandler...  " + sess.Get("username").(string) + " logined by session successfully")
//		//	return
//		//}
//	}
//	tmpl, err := template.ParseFiles("/admin_login.html")
//	checkErr(err)
//	err = tmpl.Execute(w, nil)
//	checkErr(err)
//}
