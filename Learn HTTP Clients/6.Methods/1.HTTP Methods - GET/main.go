package main

import (
	"log"
)

func main() {
	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"
	users, err := getUsers(url)
	if err != nil {
		log.Fatal("Error getting users:", err)
	}
	logUsers(users)
}
