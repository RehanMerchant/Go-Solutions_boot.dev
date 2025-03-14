package main

import (
	"encoding/json"
	"net/http"
)

func getUsers(url string) ([]User, error) {
	fullURL := url + "?sort=experience"
	res, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var users []User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
