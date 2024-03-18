package dtos

type ValidationErrorDto struct {
	Key   string
	Field string
	Error string
	Tag   string
}
