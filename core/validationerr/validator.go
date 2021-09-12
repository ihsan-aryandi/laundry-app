package validationerr

import "fmt"

type Validation struct {
	ValidationErrors map[string]map[string]string
}

type ValidationHandler func() (msg string, contents interface{})

func NewValidation() *Validation {
	return &Validation{
		ValidationErrors: make(map[string]map[string]string),
	}
}

func (v *Validation) AddValidation(fieldName string, label string, handler ValidationHandler) interface{} {
	message, contents := handler()
	if message == "" {
		return contents
	}

	m := make(map[string]string)
	m["message"] = fmt.Sprintf(message, label)
	m["format"] = message

	v.ValidationErrors[fieldName] = m
	return contents
}

func (v *Validation) Errors() map[string]map[string]string {
	if len(v.ValidationErrors) == 0 {
		return nil
	}
	return v.ValidationErrors
}