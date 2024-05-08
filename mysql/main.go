package main

import "gomysql/users"

type User struct {
	Name string `json:"name"`
}

func main() {

	alex := users.User{
		FirstName:  "Alex",
		MiddleName: "Jr.",
		LastName:   "Doe",
		Age:        30,
	}

	alex.CreateUser()
}
