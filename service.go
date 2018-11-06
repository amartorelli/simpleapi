package main

import (
	"errors"
	"strings"
)

var (
	// ErrEmpty is when an empty string is provided
	ErrEmpty = errors.New("Empty string")
)

// StringService provides a service to convert strings
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// interface implementation
type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}
