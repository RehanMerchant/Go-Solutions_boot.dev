package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func getUserById(baseURL, id string) (User, error) {
	fullURL := baseURL + "/" + id
	if err := errIfNotHTTPS(fullURL); err != nil {
		return User{}, err
	}

	res, err := http.Get(fullURL)
	if err != nil {
		return User{}, err
	}

	var user User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func logUser(user User) {
	fmt.Printf("User Info - UUID: %s, Name: %s, Role: %s, Experience: %d, Remote: %v\n",
		user.ID, user.User.Name, user.Role, user.Experience, user.Remote)
}
