package errors

import "fmt"

type ObjectNotFound struct {
	ID string
}

func (e *ObjectNotFound) Error() string {
	return fmt.Sprintf("object not found (id=%q)", e.ID)
}
