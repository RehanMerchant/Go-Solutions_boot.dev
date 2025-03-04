package main

import (
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i := range messages {
		messages[i].tags = tagger(messages[i])
	}
	return messages
}

func tagger(msg sms) []string {

	tags := []string{}

	lowerContent := strings.ToLower(msg.content)

	if strings.Contains(lowerContent, "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(lowerContent, "sale") {
		tags = append(tags, "Promo")
	}

	return tags

}
