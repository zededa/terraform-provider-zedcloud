package errors

import "fmt"

type ObjectNotFound struct {
	ID string
}

func (e *ObjectNotFound) Error() string {
	return fmt.Sprintf("object not found (id=%q)", e.ID)
}

func GetErrMsgPrefix(name, id, objectType, action string) string {
	return fmt.Sprintf("[ERROR] %s %s ( id: %s) %s Failed.", objectType, name, id, action)
}
