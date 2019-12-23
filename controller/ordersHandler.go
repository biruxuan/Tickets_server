package controller

import (
	"Tickets_server/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

//获取全部订单
func GetAllOrderss(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		//http.Error(w, "error", 404)
		w.WriteHeader(404)
	} else {
		orders, err := model.GetAllOrders()
		if err != nil {
			//t := template.Must(template.ParseFiles("views/pages/404.html"))
			//t.Execute(w, "")
			http.Error(w, err.Error(), 404)
		} else {
			//t := template.Must(template.ParseFiles("views/pages/tickets_manger.html"))
			//t.Execute(w, tickets)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(orders)
			//_, _ = fmt.Fprint(w, tickets)
		}
	}
}

//删除一张订单
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	//获取要修改的班次
	order_id := r.PostFormValue("ticket_id")
	//进行类型转换
	iOrderID, _ := strconv.ParseInt(order_id, 10, 0)
	err := model.DeleteTicketByID(iOrderID)
	if err != nil {
		w.WriteHeader(404)
		//panic(err)
	} else {
		w.WriteHeader(200)
	}
}

//增加一个订单
func AddOrder(w http.ResponseWriter, r *http.Request) {
	//ticketID := r.PostFormValue("Ticket_id")
	//trainID := r.PostFormValue("Train_id")
	//departureDate := r.PostFormValue("Departure_Date")
	//departureTime := r.PostFormValue("Departure_time")
	//startPoint := r.PostFormValue("start_point")
	//endPoint := r.PostFormValue("End_point")
	//travelTime := r.PostFormValue("Travel_time")
	//ratedLoad := r.PostFormValue("Rated_load")
	//bookedNum := r.PostFormValue("Booked_num")
	////类型转换
	//iticketID, _ := strconv.ParseInt(ticketID, 10, 0)
	//iratedLoad, _ := strconv.ParseInt(ratedLoad, 10, 0)
	//ibookedNun, _ := strconv.ParseInt(bookedNum, 10, 0)
	//ftravelTime, _ := strconv.ParseFloat(travelTime, 32)
	//
	//ticket := model.Ticket{
	//	Ticket_id:      iticketID,
	//	Train_id:       trainID,
	//	Departure_date: departureDate,
	//	Departure_time: departureTime,
	//	Start_point:    startPoint,
	//	End_point:      endPoint,
	//	Travel_time:    ftravelTime,
	//	Rated_load:     iratedLoad,
	//	Booked_num:     ibookedNun,
	//}
	//err := ticket.AddTickets()
	//if err != nil {
	//	//fmt.Fprint(w, 404)
	//	w.WriteHeader(404)
	//
	//}
	//w.WriteHeader(200)
	_ = r.ParseForm()
	if r.Method == "POST" {
		//formTicket:=make(map[string]interface{})
		t := model.Ticket{}
		_ = json.NewDecoder(r.Body).Decode(t)
		err := t.AddTickets()
		if err != nil {
			fmt.Println("订单插入数据库错误")
		}
	} else {
		w.WriteHeader(404)
	}
}
