package main

import (
	"fmt"
)

type User struct {
	Role       string `json:"role"`
	ID         string `json:"id"`
	Experience int    `json:"experience"`
	Remote     bool   `json:"remote"`
	User       struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Age      int    `json:"age"`
	} `json:"user"`
}

func logUser(user User) {
	fmt.Printf("User Name: %s, Role: %s, Experience: %d, Remote: %v\n",
		user.User.Name, user.Role, user.Experience, user.Remote)
}
