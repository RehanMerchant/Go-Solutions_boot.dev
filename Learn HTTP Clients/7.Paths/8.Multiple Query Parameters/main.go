package main

import (
	"fmt"
	"net/url"
)

func fetchTasks(baseURL, availability string) []Issue {

	limit := 5
	if availability == "Low" {
		limit = 1
	} else if availability == "Medium" {
		limit = 3
	}

	params := url.Values{}
	params.Add("sort", "estimate")
	params.Add("limit", fmt.Sprintf("%d", limit))
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	return getIssues(fullURL)
}
