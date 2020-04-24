package utils

import "regexp"

// ValidateEmail ...
func ValidateEmail(email string) bool {
	var emailRegx = regexp.MustCompile("^[a-zA-Z0-9.-_]+@[a-zA-Z0-9-]+.[a-zA-Z]{2,5}$")

	return emailRegx.MatchString(email)
}
