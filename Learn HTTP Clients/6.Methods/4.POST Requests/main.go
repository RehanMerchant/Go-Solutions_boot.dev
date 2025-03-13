package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

func main() {
	userToCreate := User{
		Role:       "Junior Developer",
		Experience: 2,
		Remote:     true,
		User: struct {
			Name     string `json:"name"`
			Location string `json:"location"`
			Age      int    `json:"age"`
		}{
			Name:     "Dan",
			Location: "NOR",
			Age:      29,
		},
	}

	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"
	apiKey := generateKey()

	fmt.Println("Retrieving user data...")
	userDataFirst, err := getUsers(url, apiKey)
	if err != nil {
		fmt.Println("Error retrieving users:", err)
		return
	}
	logUsers(userDataFirst)
	fmt.Println("---")

	fmt.Println("Creating new user...")
	creationResponse, err := createUser(url, apiKey, userToCreate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	jsonData, _ := json.Marshal(creationResponse)
	fmt.Printf("Creation response body: %s\n", string(jsonData))
	fmt.Println("---")

	fmt.Println("Retrieving user data...")
	userDataSecond, err := getUsers(url, apiKey)
	if err != nil {
		fmt.Println("Error retrieving users:", err)
		return
	}
	logUsers(userDataSecond)
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
