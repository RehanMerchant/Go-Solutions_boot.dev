package main

import "unicode"

func isValidPassword(password string) bool {

	if len(password) < 5 || len(password) > 12 {
		return false
	}

	upper := false
	digit := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			upper = true
		}
		if unicode.IsDigit(char) {
			digit = true
		}

		if upper && digit {
			return true
		}

	}

	return false

}
