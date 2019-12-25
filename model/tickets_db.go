package model

import (
	"Tickets_server/utils"
	"fmt"
)

type Ticket struct {
	Ticket_id      int64
	Train_id       string
	Departure_date string
	Departure_time string
	Start_point    string
	End_point      string
	Travel_time    float64
	Rated_load     int64
	Booked_num     int64
}

//增加一张车票
func (ticket *Ticket) AddTickets() error {
	//sql语句
	sqlStr := "insert into tickets_info(ticket_id,train_id,departure_date,departure_time,Start_point,end_point,travel_time," +
		"rated_load,booked_num) values(?,?,?,?,?,?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, ticket.Ticket_id, ticket.Train_id, ticket.Departure_date, ticket.Departure_time, ticket.Start_point, ticket.End_point,
		ticket.Travel_time, ticket.Rated_load, ticket.Booked_num)
	if err != nil {
		fmt.Println("数据库插入错误", err.Error())
		return err
	}
	return nil
}

//通过车次查询车票
func (ticket *Ticket) GetTicketByID() (*Ticket, error) {
	sqlStr := "select * from tickets_info where ticket_id=?"
	row := utils.Db.QueryRow(sqlStr, ticket.Ticket_id)
	var (
		ticket_id      int64
		train_id       string
		departure_date string
		departure_time string
		start_point    string
		end_point      string
		travel_time    float64
		rated_load     int64
		booked_num     int64
	)
	err := row.Scan(&ticket_id, &train_id, &departure_date, &departure_time, &start_point, &end_point, &travel_time,
		&rated_load, &booked_num)
	if err != nil {
		fmt.Println("通过ID查询错误")
		return nil, err
	}

	t := &Ticket{
		Ticket_id:      ticket_id,
		Train_id:       train_id,
		Departure_date: departure_date,
		Departure_time: departure_time,
		Start_point:    start_point,
		End_point:      end_point,
		Travel_time:    travel_time,
		Rated_load:     rated_load,
		Booked_num:     booked_num,
	}
	return t, nil
}

//获取全部车票
func GetAllTickets() ([]*Ticket, error) {
	//func GetAllTickets(str string) ([]*Ticket, error) {
	sqlStr := "select * from tickets_info where 1"
	//sqlStr := "select * from tickets_info where "+str

	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	var tickets []*Ticket
	for rows.Next() {
		var (
			ticketId      int64
			trainId       string
			departureDate string
			departureTime string
			startPoint    string
			endPoint      string
			travelTime    float64
			ratedLoad     int64
			bookedNum     int64
		)
		err2 := rows.Scan(&ticketId, &trainId, &departureDate, &departureTime, &startPoint, &endPoint, &travelTime,
			&ratedLoad, &bookedNum)
		if err2 != nil {
			return nil, err2
		}
		t := &Ticket{
			Ticket_id:      ticketId,
			Train_id:       trainId,
			Departure_date: departureDate,
			Departure_time: departureTime,
			Start_point:    startPoint,
			End_point:      endPoint,
			Travel_time:    travelTime,
			Rated_load:     ratedLoad,
			Booked_num:     bookedNum,
		}
		tickets = append(tickets, t)
	}
	return tickets, nil
}

//删除一张车票
func DeleteTicketByID(ID int64) error {
	sqlStr := "delete from tickets_info where ticket_id = ?"
	_, err := utils.Db.Exec(sqlStr, ID)
	if err != nil {
		return err
	}
	return nil
}

//修改已定车票
func UpdateTicketBookedNum(ID int64, status string) (int64, error) {
	var sqlStr string
	if status == "refund" {
		//退票
		sqlStr = "update tickets_info set booked_num = booked_num-1  where ticket_id =?"
	} else if status == "sold" {
		//售票
		sqlStr = "update tickets_info set booked_num = booked_num+1  where ticket_id =?"
	}

	_, err := utils.Db.Exec(sqlStr, ID)
	if err != nil {
		return -1,err
	}

	sqlQueryrow := "select booked_num from tickets_info where ticket_id = ?"
	row := utils.Db.QueryRow(sqlQueryrow, ID)
	var b_n int64
	err2:=row.Scan(&b_n)
	if err2 != nil {
		return -1,err2
	}
	//num:=make(map[string]int64)
	//num["Booked_num"]=b_n
	return b_n, nil
}
