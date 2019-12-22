package model

import (
	"Tickets_server/utils"
	"fmt"
)

type Ticket struct {
	ticket_id      int
	departure_time string
	start_point    string
	end_point      string
	travel_time    float32
	rated_load     int
	booked_num     int
}


func (ticket *Ticket) AddTickets() error {
	//sql语句
	sqlStr := "insert into tickets_info(ticket_id,departure_time,start_point,end_point,travel_time," +
		"rated_load,booked_num) values(?,?,?,?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, ticket.ticket_id, ticket.departure_time, ticket.start_point, ticket.end_point,
		ticket.travel_time, ticket.rated_load, ticket.booked_num)
	if err != nil {
		fmt.Println("数据库插入错误", err.Error())
		return err
	}
	return nil
}

func (ticket *Ticket) GetTicketByID() (*Ticket, error) {
	sqlStr := "select * from tickets_info where ticket_id=?"
	row := utils.Db.QueryRow(sqlStr, ticket.ticket_id)
	var (
		ticket_id      int
		departure_time string
		start_point    string
		end_point      string
		travel_time    float32
		rated_load     int
		booked_num     int
	)
	err := row.Scan(&ticket_id, &departure_time, &start_point, &end_point, &travel_time,
		&rated_load, &booked_num)
	if err != nil {
		fmt.Println("通过ID查询错误")
		return nil, err
	}

	t := &Ticket{
		ticket_id:      ticket_id,
		departure_time: departure_time,
		start_point:    start_point,
		end_point:      end_point,
		travel_time:    travel_time,
		rated_load:     rated_load,
		booked_num:     booked_num,
	}
	return t, nil
}

func GetAllTickets() ([]*Ticket, error) {
	sqlStr := "select * from tickets_info where 1"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	var tickets []*Ticket
	for rows.Next() {
		var (
			ticket_id      int
			departure_time string
			start_point    string
			end_point      string
			travel_time    float32
			rated_load     int
			booked_num     int
		)
		err2 := rows.Scan(&ticket_id, &departure_time, &start_point, &end_point, &travel_time,
			&rated_load, &booked_num)
		if err2 != nil {
			return nil, err2
		}
		t := &Ticket{
			ticket_id:      ticket_id,
			departure_time: departure_time,
			start_point:    start_point,
			end_point:      end_point,
			travel_time:    travel_time,
			rated_load:     rated_load,
			booked_num:     booked_num,
		}
		tickets = append(tickets, t)
	}
	return tickets, nil
}
