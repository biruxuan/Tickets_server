package model

import (
	"Tickets_server/utils"
	//"fmt"
)

type Order struct {
	Order_id   int64
	Oticket_id int64
	Name       string
	Id_card    string
	Phone      int64
}

type OrderList struct {
	TicketID int64
	TrainID        string
	Departure_date string
	Departure_time string
	Start_point    string
	End_point      string
	TravelTime     int64
	Name           string
	IDCard         string
}


//增加一个订单
func (order *Order) AddOrder() (int64, error) {
	sqlStr := "insert into orders_info(order_id,oticket_id,name,id_card,phone)" +
		"values(?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, order.Order_id, order.Oticket_id, order.Name, order.Id_card, order.Phone)
	if err != nil {
		return -1, err
	}
	queryStr := "select order_id from orders_info where oticket_id=? AND id_card = ?"
	row := utils.Db.QueryRow(queryStr, order.Oticket_id, order.Id_card)
	var id int64
	_ = row.Scan(&id)
	_,_=UpdateTicketBookedNum(order.Oticket_id,"sold")
	return id, nil
}

//删除一个订单
func DeleteOrderByID(orderID int64,ticketID int64) error {
	sqlStr := "delete from orders_info where order_id = ?"
	_, err := utils.Db.Exec(sqlStr, orderID)
	_,_=UpdateTicketBookedNum(ticketID,"refund")
	if err != nil {
		return err
	}
	return nil
}

//获取全部订单
func GetAllOrders(order_id int64) (OrderList, error) {
	sqlStr := "SELECT tickets_info.ticket_id,tickets_info.train_id,tickets_info.departure_date," +
		"tickets_info.departure_time,tickets_info.start_point,tickets_info.end_point," +
		"tickets_info.travel_time,orders_info.name,orders_info.id_card from tickets_info," +
		"orders_info WHERE  orders_info.oticket_id=tickets_info.ticket_id AND order_id=?"

	//fmt.Println(order_id)
	row := utils.Db.QueryRow(sqlStr, order_id)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}

	//var ordersList []*OrderList
	//for rows.Next() {
	//	var (
	//		trainID        string
	//		departure_date string
	//		departure_time string
	//		start_point    string
	//		end_point      string
	//		travelTime     int64
	//		nAme           string
	//		idCard         string
	//	)
	//	err2 := rows.Scan(trainID, departure_date, departure_time, start_point, end_point, travelTime, nAme, idCard)
	//	if err2 != nil {
	//		fmt.Println(err)
	//
	//		return nil, err2
	//	}
	//	t := &OrderList{
	//		trainID:        trainID,
	//		departure_date: departure_date,
	//		departure_time: departure_time,
	//		start_point:    start_point,
	//		end_point:      end_point,
	//		travelTime:     travelTime,
	//		name:           nAme,
	//		idCard:         idCard,
	//	}
	//	ordersList = append(ordersList, t)
	//}
	//return ordersList, nil

	var orderList OrderList
	err:=row.Scan(&orderList.TicketID,&orderList.TrainID, &orderList.Departure_date, &orderList.Departure_time,
		&orderList.Start_point, &orderList.End_point, &orderList.TravelTime, &orderList.Name, &orderList.IDCard)
	if err!=nil{
		return orderList,err
	}
	return orderList, err
}

//通过身份证查检验订单重复
func CheckOrder(id_card string, ticketID int64) error {
	sqlStr := "select order_id from orders_info where id_card=?  and oticket_id= ?"
	row := utils.Db.QueryRow(sqlStr, id_card, ticketID)
	var flag int64
	e := row.Scan(&flag)
	if e != nil {
		return e
	} else {
		return nil
	}

}
