package customError

type InvalidEntityCreationError struct {
	EntityName string
	Field      string
	Value      interface{}
}

func (e *InvalidEntityCreationError) Error() string {
	return "Invalid " + e.EntityName + " creation: " + e.Field + " cannot be " + string(e.Value.(string))
}
