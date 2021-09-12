package validationerr

func (*Validation) Required(value string) bool {
	return value == ""
}

func (*Validation) RequiredNum(value int) bool {
	return value < 1
}

func (*Validation) Max(value string, max int) bool {
	return len(value) > max
}

func (*Validation) Min(value string, min int) bool {
	return len(value) < min
}
