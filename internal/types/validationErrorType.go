package types

type ValidationError struct {
	Key   string
	Field string
	Error string
	Tag   string
}
