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

func Datahandler(w http.ResponseWriter, r *http.Request) {
	//tickets,_:=model.GetAllTickets()
	//t:=template.Must(template.ParseFiles("index.html"))
	//w.Write(tickets)
	//t.Execute(w,tickets)
	//if r.URL.Path == "index" && r.Method == "GET" {
	//len:=r.ContentLength
	//data:=make([]byte,len)
	fmt.Println(r.URL)
	fmt.Fprintf(w, "success")
	//}

	//len := r.ContentLength
	//body := make([]byte, len)
	//r.Body.Read(body)
	////fmt.Fprint(w, "请求体中的内容", string(body))
	////fmt.Fprint(w, "method", r.Method)
	//fmt.Fprint(w, "HEADER", r.Header)

	//fmt.Fprint(w, "URL内容", r.URL)
	//
	//fmt.Println(string(body))
	//fmt.Println(w, "URL内容", r.URL)

	//fmt.Fprintf(w, r.FormValue(""))
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

//修改已售出车票
func UpdateTicketBooked(w http.ResponseWriter, r *http.Request) {
	//获取要修改的班次
	//ticketId := r.PostFormValue("ticket_id")
	//status := r.PostFormValue("status")
	//进行类型转换
	//iTicketID, _ := strconv.ParseInt(ticketId, 10, 10)
	//num,err := model.UpdateTicketBookedNum(iTicketID, status)
	num, err := model.UpdateTicketBookedNum(2, "refund")
	if err != nil {
		http.Error(w, err.Error(), 404)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(num)
}
