package vfs

import (
	"fmt"
	"regexp"
)

// Constants for input validation
const (
	MaxInputLength = 255
	ValidChars     = "^[a-zA-Z0-9_-]+$"
)

// ValidateInput validates the input string.
func ValidateInput(input string) error {
	if len(input) == 0 || len(input) > MaxInputLength {
		return fmt.Errorf("input length must be between 1 and %d characters", MaxInputLength)
	}
	if match, _ := regexp.MatchString(ValidChars, input); !match {
		return fmt.Errorf("input contains invalid characters")
	}

	return nil
}
