package main

func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}
	return formattedMessages
}
