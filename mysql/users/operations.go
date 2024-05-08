package users

import (
	"fmt"
	"gomysql/helpers"
)

type User struct {
	Id         int    `json:"id"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Age        int    `json:"age"`
}

func checkTable() {
	userTable := `CREATE TABLE IF NOT EXISTS users
(
	id INT NOT NULL AUTO_INCREMENT,
	firstName VARCHAR(50),
	middleName VARCHAR(50),
	lastName VARCHAR(50),
	age INT,
		PRIMARY KEY(id)
);`

	_, err := helpers.DB().Query(userTable)

	if err != nil {
		panic(err)
	}
}

func (u *User) CreateUser() *User {
	checkTable()

	createQuery := fmt.Sprintf(
		"INSERT INTO users (firstName, middleName, lastName, age) VALUES ('%s', '%s', '%s', %d);",
		u.FirstName,
		u.MiddleName,
		u.LastName,
		u.Age,
	)

	fmt.Println(createQuery)

	_, err := helpers.DB().Query(createQuery)

	if err != nil {
		panic(err)
	}

	return u
}
