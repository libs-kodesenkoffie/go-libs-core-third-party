package helpers

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DB() *sql.DB {
	db, err := sql.Open("mysql", "test_user:Password@123@tcp(127.0.0.1:3306)/test_db")

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err != nil {
		panic(err)
	}

	return db
}
