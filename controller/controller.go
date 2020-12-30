package controller

import (
	"ParkingLot/memory"
	"ParkingLot/session"
	"fmt"
	"log"
	"net/http"
)

var globalSessions *session.Manager

func init() {
	// 初始化存储session的内存
	memory.Init()
	var err error
	// 定义一个全局的SessionManager
	globalSessions, err = session.NewSessionManager("memory",
		"goSessionID",
		3600)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	go globalSessions.GC()
}

func StartUp() {
	log.Println("start...")
	//go RabbitMQReceive()
	Handlers()
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Println("ListenAndServe", err.Error())
	}
}

func Handlers() {
	http.Handle("/css/", http.FileServer(http.Dir("./views/static")))
	http.Handle("/font-awesome/", http.FileServer(http.Dir("./views/static")))
	http.Handle("/js/", http.FileServer(http.Dir("./views/static")))
	http.Handle("/img/", http.FileServer(http.Dir("./views/static")))
	http.Handle("/static/", http.FileServer(http.Dir("./views")))
	http.Handle("/pages/", http.FileServer(http.Dir("./views")))

	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/forgot", forgotHandler)
	http.HandleFunc("/doLogin", doLoginHandler)
	http.HandleFunc("/checkToken", checkTokenHandler)

	//http.HandleFunc("/adminLogin", controller.AdminLoginHandler)

	http.HandleFunc("/index", indexHandler)

	//http.HandleFunc("/doLogin", doLoginHandler)

	http.HandleFunc("/doRegister", doRegisterHandler)
	//http.HandleFunc("/doLogout", doLogoutHandler)

}
