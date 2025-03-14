package main

import "fmt"

func main() {
	uuid := "2f8282cb-e2f9-496f-b144-c0aa4ced56db"

	user, err := getUserById(URL, uuid)
	if err != nil {
		fmt.Println(err)
		return
	}

	logUser(user)
}
