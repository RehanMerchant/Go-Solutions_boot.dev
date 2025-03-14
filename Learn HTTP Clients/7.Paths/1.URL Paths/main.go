package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	projects := getResources("/v1/courses_rest_api/learn-http/projects")
	fmt.Println("Projects:")
	logResources(projects)
	fmt.Println(" --- ")

	issues := getResources("/v1/courses_rest_api/learn-http/issues")
	fmt.Println("Issues:")
	logResources(issues)
	fmt.Println(" --- ")

	users := getResources("/v1/courses_rest_api/learn-http/users")
	fmt.Println("Users:")
	logResources(users)
}

func logResources(resources []map[string]any) {
	for _, resource := range resources {
		jsonResource, err := json.Marshal(resource)
		if err != nil {
			fmt.Println("Error marshalling resource:", err)
			continue
		}
		fmt.Printf(" - %s\n", jsonResource)
	}
}
