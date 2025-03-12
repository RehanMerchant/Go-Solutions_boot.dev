package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const url = "https://api.boot.dev//v1/courses_rest_api/learn-http/projects"

func main() {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("error creating request: %v", err)
	}
	defer res.Body.Close()

	var projects []Project
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&projects)
	if err != nil {
		log.Fatalf("error decoding response: %v", err)
	}

	logProjects(projects)
}

type Project struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func logProjects(projects []Project) {
	for _, p := range projects {
		fmt.Printf("Project: %s, Complete: %v\n", p.Title, p.Completed)
	}
}
