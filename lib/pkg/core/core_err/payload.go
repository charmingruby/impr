package core_err

import "fmt"

type SerializationErr struct {
	Message string `json:"message"`
}

func NewSerializationErr(name string, format string) *SerializationErr {
	return &SerializationErr{
		Message: fmt.Sprintf("unable to serialize %s to %s", name, format),
	}
}

func (e *SerializationErr) Error() string {
	return e.Message
}
