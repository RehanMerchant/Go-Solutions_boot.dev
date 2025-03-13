package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id

	jsondata, err := json.Marshal(data)
	if err != nil {
		return User{}, nil
	}

	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsondata))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)

	defer res.Body.Close()

	var user User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&user)

	if err != nil {
		return User{}, nil
	}

	return user, nil

}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest("GET", fullURL, nil)

	req.Header.Set("X-API-Key", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return User{}, err
	}

	defer res.Body.Close()

	var user User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&user)

	if err != nil {
		return User{}, nil
	}

	return user, nil

}
