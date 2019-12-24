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
	if r.Method != "POST" {
		//http.Error(w, "error", 404)
		w.WriteHeader(404)
	} else {
		order_id := r.PostFormValue("order_id")
		iOrderID, _ := strconv.ParseInt(order_id, 10, 0)
		ordersList, _ := model.GetAllOrders(iOrderID)
		//for k, v := range ordersList {
		//	fmt.Printf("%v: %v", k, v)
		//}

		//t := template.Must(template.ParseFiles("views/pages/tickets_manger.html"))
		//t.Execute(w, tickets)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(ordersList)
		fmt.Println(err)

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
	if r.Method == "POST" {
		//formTicket:=make(map[string]interface{})
		name := r.PostFormValue("name")
		idCard := r.PostFormValue("id_card")
		oticketID := r.PostFormValue("oticket_id")
		phone := r.PostFormValue("phone")

		ticketID, _ := strconv.ParseInt(oticketID, 10, 0)
		iPhone, _ := strconv.ParseInt(phone, 10, 0)
		t := model.Order{
			//Order_id:   0,
			Oticket_id: ticketID,
			Name:       name,
			Id_card:    idCard,
			Phone:      iPhone,
		}
		err := model.CheckOrder(t.Id_card, t.Oticket_id)
		if err != nil {
			fmt.Println("开始插入数据库")
			ID, _ := t.AddOrder()
			fmt.Println(ID)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(ID)
		} else {
			w.WriteHeader(404)
		}
	}
	w.WriteHeader(404)
}
