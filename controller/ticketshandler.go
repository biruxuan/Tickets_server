package controller

import (
	"Tickets_server/model"
	"fmt"
	"html/template"
	"net/http"
)

func Indexhandler(w http.ResponseWriter, r *http.Request){
	t:=template.Must(template.ParseFiles("index.html"))
	t.Execute(w,"hello world")
}

func Datahandler(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w,"hello")
}

func GetAllTickets(w http.ResponseWriter, r *http.Request) {
	tickets, _ := model.GetAllTickets()
	if tickets == nil {
		t := template.Must(template.ParseFiles("views/pages/404.html"))
		t.Execute(w, "")
	} else{
		t := template.Must(template.ParseFiles("views/pages/tickets_manger.html"))
		t.Execute(w, tickets)
	}
}
