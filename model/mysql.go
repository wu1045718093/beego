package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var connect *sql.DB

func init() {
	_, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/mysql")
	if nil != err {
		fmt.Println("connect db error: ", err)
	}
}

func MysqlQuerry() {

}

func MysqlInsert() {

}
