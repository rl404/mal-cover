package errors

import (
	"errors"
	"fmt"
	"strings"
)

// Error list.
var (
	ErrInternalServer = errors.New("internal server error")
	ErrInternalCache  = errors.New("internal cache error")
)

// ErrRequiredField is error for missing field.
func ErrRequiredField(str string) error {
	return fmt.Errorf("required field %s", str)
}

// ErrGTField is error for greater than field.
func ErrGTField(str, value string) error {
	return fmt.Errorf("field %s must be greater than %s", str, value)
}

// ErrOneOfField is error for oneof field.
func ErrOneOfField(str, value string) error {
	return fmt.Errorf("field %s must be one of %s", str, strings.Join(strings.Split(value, " "), "/"))
}

// ErrStyleField is error for style field.
func ErrStyleField() error {
	return fmt.Errorf("empty style param\nplease check your list style\n\ntry this example:\n\n.animetitle[href*='/{id}/']:before{background-image:url({url})}")
}
