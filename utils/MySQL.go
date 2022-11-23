package utils

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"os/signal"
	"time"
)

var pool *sql.DB // Database connection pool.

func Connect() {
	// dsn格式: [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	// Config.formatDSN可以传入一个结构体进行解析
	dsn := "root:zyt123456@tcp(121.5.154.155:3306)/spike?charset=utf8"
	var err error

	// Opening a driver typically will not attempt to connect to the database.
	pool, err = sql.Open("mysql", dsn)
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
		<-appSignal
		stop()
	}()

	Ping(ctx)
	log.Printf("Mysql(%v) has connected", dsn)
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

func GetConn() *sql.DB {
	if pool == nil {
		Connect()
	}
	return pool
}

func Close() {
	if pool != nil {
		defer pool.Close()
	}
}
