package main

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, user := range messagedUsers {
		if _, exists := validUsers[user]; exists {
			validUsers[user]++
		}
	}
}
