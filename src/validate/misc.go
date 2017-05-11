package validate

import "strings"

func Email(email string) (bool, string) {
	if len(email) < 5 {
		return false, "Email address is too short"
	}

	atPos := strings.Index(email, "@")
	if atPos == -1 {
		return false, "Email address is missing @"
	}

	return true, ""
}
