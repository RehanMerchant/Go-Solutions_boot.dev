package main

import (
	"fmt"
)

const baseUrl = "https://api.boot.dev"

func main() {
	issues, err := getResources(baseUrl + "/v1/courses_rest_api/learn-http/issues?limit=1")
	if err != nil {
		fmt.Println("Error getting issues:", err)
		return
	}
	fmt.Println("Issue:")
	logResources(issues)
	fmt.Println("---")

	projects, err := getResources(baseUrl + "/v1/courses_rest_api/learn-http/projects?limit=1")
	if err != nil {
		fmt.Println("Error getting projects:", err)
		return
	}
	fmt.Println("Project:")
	logResources(projects)
	fmt.Println("---")

	users, err := getResources(baseUrl + "/v1/courses_rest_api/learn-http/users?limit=1")
	if err != nil {
		fmt.Println("Error getting users:", err)
		return
	}
	fmt.Println("User:")
	logResources(users)
}
