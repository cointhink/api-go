package validate

import "strings"
import "regexp"

func Email(email string) (bool, string) {
	if len(email) < 5 {
		return false, "Email address is too short"
	}

	atPos := strings.Index(email, "@")
	if atPos == -1 {
		return false, "Email address is missing @"
	}

	match, _ := regexp.MatchString("^[[:graph:]]+@[[:graph:]]+$", email)
	if !match {
		return false, "non-conforming email address"
	}

	return true, ""
}
