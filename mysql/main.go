package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
}

func main() {

	// db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test_db")
	db, err := sql.Open("mysql", "test_user:Password@123@tcp(127.0.0.1:3306)/test_db")

	// fmt.Printf("%+v", err)

	if err != nil {
		panic(err)
	}

	// res, _ := db.Query("CREATE TABLE users ( name VARCHAR(255));")
	// res, _ := db.Query("insert into users values ('Alex');")

	res, _ := db.Query("select name from users;")

	for res.Next() {
		var user User
		_ = res.Scan(&user.Name)
		fmt.Printf("%+v\n", user)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
