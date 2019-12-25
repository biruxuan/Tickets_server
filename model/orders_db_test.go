package model

import (
	"fmt"
	"testing"
)

//func Testmain(t *testing.T){
//	t.Run("测试添加",TestOrder_AddOrder)
//	t.Run("测试显示",TestGetAllOrders)
//}

//func TestOrder_AddOrder(t *testing.T) {
//	order := Order{
//		//Order_id:   3,
//		Oticket_id: 12,
//		Name:       "张三",
//		Id_card:    "37232319990101123X",
//		Phone:      12345678910,
//	}
//	err:= order.AddOrder()
//	if err!=nil{
//		fmt.Println(err)
//	}
//}

//func TestDeleteOrderByID(t *testing.T) {
//	_=DeleteOrderByID(4)
//}

func TestGetAllOrders(t *testing.T) {
	orders,_ := GetAllOrders(56)

		fmt.Println(orders)

}

//func TestCheckOrder(t *testing.T) {
//	id_card := "123465789"
//	_ = CheckOrder(id_card, 1)
//}
