package main

import (
	"Tickets_server/controller"
	"net/http"
)

func main(){
	//静态资源管理
	//http.Handle("/static",http.)
	//增加车票
	//http.HandleFunc()
	//购买车票
	//退票
	//查看全部车票
	http.HandleFunc("/index",controller.Indexhandler)
	http.Handle("/main.js/",http.StripPrefix("/main.js/",http.FileServer(http.Dir("view/static"))))
	http.HandleFunc("/data",controller.Datahandler)
	http.HandleFunc("/alltickets",controller.GetAllTickets)
	http.ListenAndServe(":8080",nil)
}