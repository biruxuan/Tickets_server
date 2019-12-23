package model

import (
	"fmt"
	"testing"
)

//func TestMain(m *testing.M){
//	ticket:=&Ticket{
//		ticket_id:1,
//		start_point:"青岛",
//		end_point:"济南",
//		departure_time:"2019-12-22 12:00:00",
//		travel_time:3.5,
//		rated_load:50,
//		booked_num:15,
//	}
//	m.Run("TestTicket_GetTicketByID",TestTicket_GetTicketByID)
//}

//func TestTicket_AddTickets(t *testing.T) {
//	ticket := &Ticket{
//		//Ticket_id:      4,
//		Train_id:       "Z104",
//		Departure_date: "2019-12-24 ",
//		Departure_time: "16:20:00",
//		Start_point:    "青岛",
//		End_point:      "北京",
//		Travel_time:    5,
//		Rated_load:     50,
//		Booked_num:     15,
//	}
//	err := ticket.AddTickets()
//	if err != nil {
//		fmt.Println("Error")
//	}
//}

//}
//
//func TestTicket_GetTicketByID(t *testing.T) {
//	ticket := &Ticket{
//		ticket_id:      1,
//	}
//
//	t_, _ := ticket.GetTicketByID()
//	fmt.Println(t_)
//}
//
//
//func TestTicket_DeleteTicket(t *testing.T) {
//
//	DeleteTicketByID(1)
//}

//func TestUpdateTicketBookedNum(t *testing.T) {
//	err:=UpdateTicketBookedNum(1,"refund")
//	if err!=nil{
//		panic(err)
//	}
//}
//
func TestTicket_GetAllTickets(t *testing.T) {
	//ticket:=Ticket{}
	str:="train_id=\"Z103\""
	tickets, _ := GetAllTickets(str)
	for k, v := range tickets {
		fmt.Printf("%v:%v", k, v)
	}
}
