package errors

import "fmt"

type ObjectNotFound struct {
	ID string
}

func (e *ObjectNotFound) Error() string {
	return fmt.Sprintf("object not found (id=%s)", e.ID)
}

type Error struct {
	ID         string
	objectType string
	objectName string
	action     string
	err        error
}

func New(id, objectType, objectName, action string, err error) *Error {
	return &Error{
		ID:         id,
		objectType: objectType,
		objectName: objectName,
		action:     action,
		err:        err,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf(
		"[ERROR] could not %s %s %s (id=%s): %s",
		e.action,
		e.objectType,
		e.objectName,
		e.ID,
		e.err.Error(),
	)
}
