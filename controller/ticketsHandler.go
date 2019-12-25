package controller

import (
	"Tickets_server/model"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Indexhandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, "")
}

func AddTicket(w http.ResponseWriter, r *http.Request) {
	//ticket := model.Ticket{
	//	Ticket_id:      r.PostFormValue("ticket_id"),
	//	Departure_time: r.PostFormValue("departure_time"),
	//	Start_point:    r.PostFormValue("start_point"),
	//	End_point:      r.PostFormValue("end_point"),
	//	Travel_time:    r.PostFormValue("travel_time"),
	//	Rated_load:     r.PostFormValue("rated_load"),
	//	Booked_num:     r.PostFormValue("booked_num"),
	//}
	////类型转换
	//iticketID, _ := strconv.ParseInt(ticketID, 10, 0)
	//iratedLoad, _ := strconv.ParseInt(ratedLoad, 10, 0)
	//ibookedNun, _ := strconv.ParseInt(bookedNum, 10, 0)
	//ftravelTime, _ := strconv.ParseFloat(travelTime, 32)
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(404)
	}
	if r.Method == "POST" {
		//初始化请求变量结构
		//formTickets:=make(map[string]interface{})
		length := r.ContentLength
		formTickets := make([]model.Ticket, length)
		//调用json包的解析，解析请求Body
		err = json.NewDecoder(r.Body).Decode(&formTickets)
		if err != nil {
			fmt.Println("车票json解析错误", err.Error())
		}
		for _, v := range formTickets {
			ticket := v
			err = ticket.AddTickets()
			if err != nil {
				log.Println("批量写入数据库错误", err)
			}
		}
	} else {
		w.WriteHeader(404)
	}
}

//获取全部车票
func GetAllTickets(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tickets, err := model.GetAllTickets()
		if err != nil {
			http.Error(w, err.Error(), 404)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(tickets)
	} else {
		w.WriteHeader(404)
	}
}



//删除一张车票
func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	//获取要修改的班次
	ticketID := r.PostFormValue("ticket_id")
	//进行类型转换
	iTicketID, _ := strconv.ParseInt(ticketID, 10, 0)
	err := model.DeleteTicketByID(iTicketID)
	if err != nil {
		http.Error(w, err.Error(), 404)
		//panic(err)
	} else {
		_, _ = fmt.Fprint(w, "true")
	}
}