package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

func getMessageText(analytics *Analytics, message Message) {

	analytics.MessagesTotal++

	if message.Success {
		*&analytics.MessagesSucceeded++
	} else {
		*&analytics.MessagesFailed++
	}

}
