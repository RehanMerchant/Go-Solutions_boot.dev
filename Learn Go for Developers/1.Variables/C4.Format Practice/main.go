package main

import "fmt"

func main() {
	fname := "Dalinar"
	lname := " Kholin"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."

	// Don't touch above this line

	userLog := fmt.Sprintf("Name: %s, Age: %d, Rate: %.1f, Is Subscribed: %t, Message: %s", fname+lname, age, messageRate, isSubscribed, message)

	// Don't touch below this line

	fmt.Println(userLog)
}
