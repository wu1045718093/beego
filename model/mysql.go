package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var pool *sql.DB // Database connection pool.

type login_user struct {
	user_id   int
	user_name string
	password  string
}

func init() {
	var err error

	// Opening a driver typically will not attempt to connect to the database.
	pool, err = sql.Open("mysql", "root:12345678@tcp(localhost:3306)/login_user")
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("unable to use data source name", err)
	}
	pool.SetMaxIdleConns(10)
	pool.SetMaxOpenConns(10)

	if err := pool.Ping(); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

}

func PrepareQuerryLogin(user_name string, password string) (result bool) {
	var u login_user
	sqlStr := "select user_id, user_name, password from user where (user_name=? AND password=?)"
	stmt, err := pool.Prepare(sqlStr)
	if err != nil {
		fmt.Println("Prepare failed, err:%v", err)
		return false
	}
	defer stmt.Close()

	err = stmt.QueryRow(user_name, password).Scan(&u.user_id, &u.user_name, &u.password)
	if err != nil {
		fmt.Println("querry failed, err:%v", err)
		return false
	}

	return true
}

func PrepareQuerryRegister(user_name string) (result bool, err error) {
	var u login_user
	sqlStr := "select user_id, user_name, password from user where (user_name=?)"
	stmt, err := pool.Prepare(sqlStr)
	if err != nil {
		fmt.Println("Prepare failed, err:%v", err)
		return false, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(user_name).Scan(&u.user_id, &u.user_name, &u.password)
	if err != nil {
		fmt.Println("querry failed, err:%v", err)
		return true, err
	}

	return false, err
}

func PrepareInsert(user string, password string) (result int) {
	ok, err := PrepareQuerryRegister(user)
	if ok != true {
		fmt.Println(" user exit, err:%v", err)
		return -1
	}

	sqlStr := "Insert into user(user_name, password) values(?,?)"
	stmt, err := pool.Prepare(sqlStr)
	if err != nil {
		fmt.Println("Prepare failed, err:%v", err)
		return -2
	}

	defer stmt.Close()

	_, err = stmt.Exec(user, password)
	if err != nil {
		fmt.Println("Insert failed, err:%v", err)
		return -3
	}

	return 0
}
