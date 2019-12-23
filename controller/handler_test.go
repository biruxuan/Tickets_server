package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteOrder(t *testing.T) {
	req,err:=http.NewRequest("GET","/deleteOrder")
	req.Body:="{\"Order_id\":\"3\"}"
	//创建 ResponseRecorder 记录响应
	rr:=httptest.NewRecorder()
	DeleteOrder(rr,req)
}