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
	//formData=
	//json.NewDecoder(r.Body).Decode(&formData)
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
	//name := r.PostFormValue("name")
	//idCard := r.PostFormValue("id_card")
	//phone := r.PostFormValue("phone")
	//fmt.Println(r.Form["id_card"])
	//fmt.Println(r.Form["phone"])
	//ticketID := r.PostFormValue("Ticket_id")
	//name := r.PostFormValue("name")
	//id_card := r.PostFormValue("id_card")
	//phone := r.PostFormValue("phone")
	//类型转换
	//fmt.Println(name)
	//fmt.Println(idCard)
	//fmt.Println(phone)

	//iPhone, _ := strconv.ParseInt(phone, 10, 0)
	//iratedLoad, _ := strconv.ParseInt(ratedLoad, 10, 0)
	//ibookedNun, _ := strconv.ParseInt(bookedNum, 10, 0)
	//ftravelTime, _ := strconv.ParseFloat(travelTime, 32)

	//order := model.Order{
	//	//Order_id:   13,
	//	Oticket_id: 14,
	//	Name:       name,
	//	Id_card:    idCard,
	//	Phone:      iPhone,
	//}
	//err := order.AddOrder()
	//if err != nil {
	//	//fmt.Fprint(w, 404)
	//	w.WriteHeader(404)
	//
	//}
	//w.WriteHeader(200)

	if r.Method == "POST" {
		//formTicket:=make(map[string]interface{})
		name:= r.PostFormValue("name")
		idCard:=r.PostFormValue("id_card")
		oticketID:=r.PostFormValue("oticket_id")
		phone:=r.PostFormValue("phone")

		//len:=r.ContentLength
		//formData:=make(map[string]string,len)
		//_ = json.NewDecoder(r.Body).Decode(formData)
		ticketID,_:=strconv.ParseInt(oticketID,10,0)
		iPhone,_:=strconv.ParseInt(phone,10,0)
		t:=model.Order{
			//Order_id:   0,
			Oticket_id: ticketID,
			Name:       name,
			Id_card:    idCard,
			Phone:      iPhone,
		}
		fmt.Println(t.Name)
		fmt.Println(t.Id_card)
		fmt.Println(t.Phone)
		fmt.Println(t.Oticket_id)


		err := t.AddOrder()
		if err != nil {
			fmt.Println("订单插入数据库错误")
		}
	} else {
		w.WriteHeader(200)
	}
}
