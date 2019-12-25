package controller

import (
	"Tickets_server/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

//根据订单号获取一个订单
func GetAllOrderss(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
	} else {
		orderID := r.PostFormValue("order_id")
		iOrderID, _ := strconv.ParseInt(orderID, 10, 0)
		orderList, err := model.GetAllOrders(iOrderID)
		//fmt.Println(orderList)
		if err != nil {
			w.WriteHeader(404)

		} else {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(orderList)
		}
	}
}

//删除一张订单
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	//获取要修改的班次
	orderID := r.PostFormValue("order_id")
	ticketID := r.PostFormValue("ticket_id")

	iOrderID, _ := strconv.ParseInt(orderID, 10, 0)
	iticketID, _ := strconv.ParseInt(ticketID, 10, 0)

	fmt.Println(iOrderID)
	err := model.DeleteOrderByID(iOrderID,iticketID)
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
			ID,_ := t.AddOrder()
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(ID)
		} else {
			w.WriteHeader(404)
		}
	}
	w.WriteHeader(404)
}
