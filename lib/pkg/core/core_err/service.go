package core_err

import "fmt"

type ResourceNotFoundErr struct {
	Message string `json:"message"`
}

func NewResourceNotFoundErr(resource string) *ResourceNotFoundErr {
	return &ResourceNotFoundErr{
		Message: fmt.Sprintf("%s not found", resource),
	}
}

func (e *ResourceNotFoundErr) Error() string {
	return e.Message
}

type InvalidFieldFormatErr struct {
	Message string `json:"message"`
}

func NewInvalidFieldFormatErr(field string, originalErr error) *InvalidFieldFormatErr {
	return &InvalidFieldFormatErr{
		Message: fmt.Sprintf("invalid %s format: %s", field, originalErr.Error()),
	}
}

func (e *InvalidFieldFormatErr) Error() string {
	return e.Message
}

type ConflictErr struct {
	Message string `json:"message"`
}

func NewConflictErr(field string) *ConflictErr {
	return &ConflictErr{
		Message: fmt.Sprintf("%s is already taken", field),
	}
}

func (e *ConflictErr) Error() string {
	return e.Message
}
