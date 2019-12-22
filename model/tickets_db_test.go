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
//		ticket_id:      2,
//		start_point:    "青岛",
//		end_point:      "济南",
//		departure_time: "2019-12-23 14:00:00",
//		travel_time:    3,
//		rated_load:     40,
//		booked_num:     25,
//	}
//	err := ticket.AddTickets()
//	if err != nil {
//		fmt.Println("Error")
//	}
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

func TestTicket_GetAllTickets(t *testing.T) {
	//ticket:=Ticket{}
	tickets,_:=GetAllTickets()
	for k,v:=range tickets{
		fmt.Printf("%v:%v",k,v)
	}
}