package main

import (
	"log"
)

func main() {
	baseURL := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"

	users, err := getUsers(baseURL)
	if err != nil {
		log.Fatal(err)
	}
	logUsers(users)

}
