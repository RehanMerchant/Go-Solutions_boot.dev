package main

import (
	"fmt"
)

func getMailtoLinkForEmail(email string) string {

	return fmt.Sprintf("mailto:%s", email)

}
