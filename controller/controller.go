package controller

import (
	"ParkingLot/memory"
	"ParkingLot/session"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	globalSessions *session.Manager
	Info           *log.Logger
	Error          *log.Logger
)

func init() {
	// 初始化存储session的内存
	memory.Init()
	var err error
	// 定义一个全局的SessionManager
	globalSessions, err = session.NewSessionManager("memory",
		"goSessionID",
		3600)
	if err != nil {
		fmt.Println(err)
		return
	}
	go globalSessions.GC()

	// 初始化log
	f, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatalf("file open error : %v", err)
	}
	Info = log.New(f, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(f, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// 错误类型日志
func checkErr(err error) {
	if err != nil {
		Error.Println(err)
	}
}

// 正常日志输出
func logInfo(info ...interface{}) {
	Info.Println(info)
}

func StartUp() {
	logInfo("start...")
	go RabbitMQReceive()
	Handlers()
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		logInfo("ListenAndServe", err.Error())
	}
}

func Handlers() {
	http.Handle("/css/", http.FileServer(http.Dir("./views/static")))
	http.Handle("/admin/", http.FileServer(http.Dir("./views/static")))
	http.Handle("/font/", http.FileServer(http.Dir("./views/static")))
	http.Handle("/js/", http.FileServer(http.Dir("./views/static")))
	http.Handle("/img/", http.FileServer(http.Dir("./views/static")))
	http.Handle("/static/", http.FileServer(http.Dir("./views")))
	http.Handle("/pages/", http.FileServer(http.Dir("./views")))

	// static
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/index", indexHandler)

	http.HandleFunc("/doLogin", doLoginHandler)
	http.HandleFunc("/doRegister", doRegisterHandler)
	http.HandleFunc("/doLogout", doLogoutHandler)
	http.HandleFunc("/checkToken", checkTokenHandler)
	http.HandleFunc("/getCaptcha", getCaptchaHandler)
	http.HandleFunc("/getSpot", getSpotHandler)
	http.HandleFunc("/booking", bookingHandler)
	http.HandleFunc("/cancelBooking", cancelBookingHandler)
	http.HandleFunc("/entry", entryHandler)
	http.HandleFunc("/out", outHandler)
	http.HandleFunc("/payFee", payFeeHandler)

	http.HandleFunc("/adminLogin", adminLoginHandler)
	http.HandleFunc("/adminLogout", adminLogoutHandler)
	http.HandleFunc("/doAdminLogin", doAdminLoginHandler)
	http.HandleFunc("/adminUsersDetail", adminUsersDetailHandler)
	http.HandleFunc("/adminSpotsDetail", adminSpotsDetailHandler)
	http.HandleFunc("/adminBookingsDetail", adminBookingsDetailHandler)
}
