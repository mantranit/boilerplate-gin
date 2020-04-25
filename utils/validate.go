package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Validate ...
var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

// ValidateEmail ...
func ValidateEmail(email string) bool {
	var emailRegx = regexp.MustCompile("^[a-zA-Z0-9.-_]+@[a-zA-Z0-9-]+.[a-zA-Z]{2,5}$")

	return emailRegx.MatchString(email)
}
