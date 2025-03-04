package main

import (
	"strings"
)

func countDistinctWords(messages []string) int {

	wordMap := make(map[string]bool)

	for _, msg := range messages {

		words := strings.Fields(strings.ToLower(msg))

		for _, word := range words {
			wordMap[word] = true
		}

	}
	return len(wordMap)
}
