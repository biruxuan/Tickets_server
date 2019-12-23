package main

import (
	"Tickets_server/controller"
	"net/http"
)

func main() {
	//静态资源管理
	http.Handle("/view/",http.StripPrefix("/view/",http.FileServer(http.Dir("./view/"))))
	//增加车票
	//http.HandleFunc()
	//购买车票
	//退票
	//查看全部车票
	http.HandleFunc("/index", controller.Indexhandler)
	http.HandleFunc("/", controller.Indexhandler)

	//http.HandleFunc("/index/", controller.Datahandler)

	http.HandleFunc("/addticket", controller.AddTicket)
	http.HandleFunc("/alltickets", controller.GetAllTickets)
	http.HandleFunc("/update", controller.UpdateTicketBooked)
	http.HandleFunc("/delete", controller.DeleteTicket)

	//h:=http.FileServer(http.Dir("./view/static"))
	http.ListenAndServe(":8080", nil)
}
