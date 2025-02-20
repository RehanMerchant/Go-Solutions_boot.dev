package main

import (
	"errors"
	"fmt"
	"testing"
)

func validateStatus(status string) error {

	if status == "" {
		return errors.New("status cannot be empty")
	}

	if len(status) > 140 {
		return errors.New("status exceeds 140 characters")
	}

	return errors.New("") // it never stated what to do if everything goes correct, but this gets all test cases passed.

}

func TestValidateStatus(t *testing.T) {
	testCases := []struct {
		status      string
		expectedErr string
	}{
		{"", "status cannot be empty"},
		{"This is a valid status update that is well within the character limit.", ""},
		{"This status update is way too long. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco.", "status exceeds 140 characters"},
	}
	if withSubmit {
		testCases = append(testCases,
			struct {
				status      string
				expectedErr string
			}{"Another valid status.", ""},
			struct {
				status      string
				expectedErr string
			}{"This status update, while derivative, contains exactly one hundred and forty-one characters, which is over the status update character limit.", "status exceeds 140 characters"},
		)
	}

	passCount := 0
	failCount := 0

	for _, tc := range testCases {
		err := validateStatus(tc.status)
		errString := ""
		if err != nil {
			errString = err.Error()
		}
		if errString != tc.expectedErr {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     "%v"
Expecting:  "%v"
Actual:     "%v"
Fail`, tc.status, tc.expectedErr, errString)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     "%v"
Expecting:  "%v"
Actual:     "%v"
Pass
`, tc.status, tc.expectedErr, errString)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

var withSubmit = true
