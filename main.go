package main

import (
	"Tickets_server/controller"
	"net/http"
)

func main() {
	//静态资源管理
	http.Handle("/view/",http.StripPrefix("/view/",http.FileServer(http.Dir("./view/"))))
//车票部分
	http.HandleFunc("/index", controller.Indexhandler)
	//http.HandleFunc("/", controller.Indexhandler)
	http.HandleFunc("/addTicket", controller.AddTicket)
	http.HandleFunc("/allTickets", controller.GetAllTickets)
	http.HandleFunc("/delete", controller.DeleteTicket)
//订单部分
	http.HandleFunc("/buyticket", controller.AddOrder)
	http.HandleFunc("/refundorder", controller.DeleteOrder)
	http.HandleFunc("/queryticket", controller.GetAllOrderss)
	//http.HandleFunc("/update", controller.UpdateTicketBooked)

	//h:=http.FileServer(http.Dir("./view/static"))
	http.ListenAndServe(":8080", nil)
}
