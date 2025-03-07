package main

func addEmailsToQueue(emails []string) chan string {

	length := len(emails)
	ch := make(chan string, length)

	for _, val := range emails {
		ch <- val
	}

	return ch

}
