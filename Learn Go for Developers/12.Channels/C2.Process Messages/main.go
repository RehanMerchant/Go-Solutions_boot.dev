package main

import (
	"time"
)

func processMessages(messages []string) []string {
	results := make(chan string, len(messages)) // Buffered channel to store results
	for _, msg := range messages {
		go func(m string) {
			results <- process(m)
		}(msg)
	}

	processedMessages := make([]string, len(messages))
	for i := 0; i < len(messages); i++ {
		processedMessages[i] = <-results
	}

	return processedMessages
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}
