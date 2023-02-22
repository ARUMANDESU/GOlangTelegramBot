package utils

import (
	"errors"
	"strconv"
	"strings"
)

func IsPercentage(s string) (float64, error) {
	// Remove any whitespace and trailing percent symbol
	s = strings.TrimSpace(s)
	s = strings.TrimSuffix(s, "%")
	InvalidError := errors.New("Invalid percentage input!")

	// Attempt to parse as a float
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		// Verify that the value is between 0 and 100
		if f >= 0 && f <= 100 {
			return f, nil
		}
	}

	// Value is not a valid percentage
	return 0, InvalidError
}
