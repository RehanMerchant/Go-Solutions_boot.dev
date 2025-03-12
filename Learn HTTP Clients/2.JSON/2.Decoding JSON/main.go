package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getIssues(url string) ([]Issue, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var issues []Issue
	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&issues); err != nil {
		fmt.Println("error decoding response body")
		return nil, fmt.Errorf("error decoding request: %w", err)
	}

	return issues, nil

}
