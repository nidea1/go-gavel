package common

type SingleValidationError struct {
	Field   string
	Code    string
	Message string
}

type ValidationErrors struct {
	Errors []SingleValidationError
}
