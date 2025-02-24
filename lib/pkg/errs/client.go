package errs

type ClientUncaughtErr struct {
	Message string `json:"message"`
}

func NewClientUncaughtErr() *ClientUncaughtErr {
	return &ClientUncaughtErr{
		Message: "client uncaught error",
	}
}

func (e *ClientUncaughtErr) Error() string {
	return e.Message
}
