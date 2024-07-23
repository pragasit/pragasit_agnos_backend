package main

import "unicode"

func calculateSteps(password string) int {
	n := len(password)
	hasLower, hasUpper, hasDigit := false, false, false

	for i := 0; i < n; {
		if unicode.IsLower(rune(password[i])) {
			hasLower = true
		}
		if unicode.IsUpper(rune(password[i])) {
			hasUpper = true
		}
		if unicode.IsDigit(rune(password[i])) {
			hasDigit = true
		}
	}
	return n
}
