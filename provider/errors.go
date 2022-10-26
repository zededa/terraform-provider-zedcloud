package provider

import "fmt"

type ObjectNotFound struct {
	id string
}

func (e *ObjectNotFound) Error() string {
	return fmt.Sprintf("object not found (id=%q)", e.id)
}
