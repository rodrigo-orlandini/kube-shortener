package customError

import "fmt"

type NotFoundError struct {
	Entity string
	Field  string
	Value  string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found with %s: %s", e.Entity, e.Field, e.Value)
}
