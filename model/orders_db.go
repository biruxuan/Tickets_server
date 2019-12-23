package model

import "Tickets_server/utils"

type Order struct {
	Order_id   int64
	Oticket_id int64
	Name       string
	Id_card    string
	Phone      int64
}

//增加一个订单
func (order *Order) AddOrder() error {
	sqlStr := "insert into orders_info(order_id,oticket_id,name,id_card,phone)" +
		"values(?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, order.Order_id, order.Oticket_id, order.Name, order.Id_card, order.Phone)
	if err != nil {
		return err
	}
	return nil
}

//删除一个订单
func DeleteOrderByID(ID int64) error {
	sqlStr := "delete from orders_info where order_id = ?"
	_, err := utils.Db.Exec(sqlStr, ID)
	if err != nil {
		return err
	}
	return nil
}


//获取全部订单
func GetAllOrders() ([]*Order, error) {
	//func GetAllTickets(str string) ([]*Ticket, error) {
	sqlStr := "select * from orders_info where 1"
	//sqlStr := "select * from tickets_info where "+str

	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	var orders []*Order
	for rows.Next() {
		var (
			orderId   int64
			oticketId int64
			nAme      string
			idCard    string
			pHone     int64
		)
		err2 := rows.Scan(&orderId, &oticketId, &nAme, &idCard, &pHone)
		if err2 != nil {
			return nil, err2
		}
		t := &Order{
			Order_id:   orderId,
			Oticket_id: oticketId,
			Name:       nAme,
			Id_card:    idCard,
			Phone:      pHone,
		}
		orders = append(orders, t)
	}
	return orders, nil
}
