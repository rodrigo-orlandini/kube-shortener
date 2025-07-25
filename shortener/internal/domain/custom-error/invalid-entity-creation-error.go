package customError

import "fmt"

type InvalidEntityCreationError struct {
	EntityName string
	Field      string
	Value      string
}

func (e *InvalidEntityCreationError) Error() string {
	return fmt.Sprintf("invalid %s creation: %s equals %s", e.EntityName, e.Field, e.Value)
}
