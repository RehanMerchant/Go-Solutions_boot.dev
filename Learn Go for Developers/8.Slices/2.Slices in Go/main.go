package main

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {

	switch plan {
	case "free":
		return messages[0:2], nil
	case "pro":
		return messages[:], nil
	default:
		return nil, errors.New("unsupported plan")
	}

}
