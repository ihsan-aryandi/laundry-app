package validationerr

import "fmt"

func RequiredMessage() string {
	return "%v is required"
}

func MinMessage(min int) string {
	m := fmt.Sprintf("%d characters length", min)
	return "%v minimum is " + m
}

func MaxMessage(max int) string {
	m := fmt.Sprintf("%d characters length", max)
	return "%v maximum is " + m
}