package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/tickets")
	if err != nil {
		//panic(err.Error())
		fmt.Println("连接tickets_info数据库失败", err.Error())
		return
	}
}
