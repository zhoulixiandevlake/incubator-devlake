package utils

import (
	"strconv"
	"time"
)

// Return defaultValue if text is empty, or try to convert it to int
func StrToIntOr(text string, defaultValue int) (int, error) {
	if text == "" {
		return defaultValue, nil
	}
	return strconv.Atoi(text)
}

// Return defaultValue if text is empty, or try to convert it to time.Duration
func StrToDurationOr(text string, defaultValue time.Duration) (time.Duration, error) {
	if text == "" {
		return defaultValue, nil
	}
	return time.ParseDuration(text)
}