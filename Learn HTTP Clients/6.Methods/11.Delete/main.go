package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	userId := "0194fdc2-fa2f-4cc0-81d3-ff12045b73c8"
	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"
	apiKey := generateKey()

	users, err := getUsers(url, apiKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Logging user records:")
	logUsers(users)
	fmt.Println("---")

	err = deleteUser(url, userId, apiKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted user with id: %s\n", userId)
	fmt.Println("---")

	newUsers, err := getUsers(url, apiKey)
	if err != nil {
		log.Fatal(err)
	}
	logUsers(newUsers)
	fmt.Println("---")
}

func generateKey() string {
	const characters = "ABCDEF0123456789"
	result := ""
	rand.New(rand.NewSource(0))
	for i := 0; i < 16; i++ {
		result += string(characters[rand.Intn(len(characters))])
	}
	return result
}
