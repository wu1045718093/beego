package model

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var pool *sql.DB // Database connection pool.

func init() {
	var err error
	// Opening a driver typically will not attempt to connect to the database.
	pool, err = sql.Open("mysql", "root:12345678@tcp(localhost:3306)/mysql")
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("unable to use data source name", err)
	}

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt)

	go func() {
		select {
		case <-appSignal:
			stop()
		}
	}()

	Ping(ctx)

	Query(ctx, 64)
}

// Ping the database to verify DSN provided by the user is valid and the
// server accessible. If the ping fails exit the program with an error.
func Ping(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}

// Query the database for the information requested and prints the results.
// If the query fails exit the program with an error.
func Query(ctx context.Context, id int64) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var name string
	err := pool.Query(ctx, "select user, password from login_user where user=teng").Scan(&)
	if err != nil {
		log.Fatal("unable to execute search query", err)
	}
	log.Println("name=", name)
}
