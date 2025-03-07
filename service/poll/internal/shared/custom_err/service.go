package custom_err

import "fmt"

type MultipleProcessErr struct {
	Message string `json:"message"`
}

type ProcessErr struct {
	Identitifer string `json:"identifier"`
	Reason      string `json:"reason"`
}

func NewMultipleProcessErr(errs []ProcessErr) *MultipleProcessErr {
	msg := "multiple errors on: "

	for idx, err := range errs {
		if idx == len(errs)-1 {
			msg += fmt.Sprintf("%s: %s.", err.Identitifer, err.Reason)
			break
		}

		msg += fmt.Sprintf("%s: %s, ", err.Identitifer, err.Reason)
	}

	return &MultipleProcessErr{
		Message: msg,
	}
}

func (e *MultipleProcessErr) Error() string {
	return e.Message
}

type InvalidActionErr struct {
	Message string `json:"message"`
}

func NewInvalidActionErr(msg string) *InvalidActionErr {
	return &InvalidActionErr{
		Message: msg,
	}
}

func (e *InvalidActionErr) Error() string {
	return e.Message
}
