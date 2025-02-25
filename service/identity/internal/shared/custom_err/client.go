package custom_err

type ClientUncaughtErr struct {
	Message string `json:"message"`
}

func NewClientUncaughtErr(err error) *ClientUncaughtErr {
	return &ClientUncaughtErr{
		Message: "client uncaught error: " + err.Error(),
	}
}

func (e *ClientUncaughtErr) Error() string {
	return e.Message
}

type UserNotConfirmedErr struct {
	Message string `json:"message"`
}

func NewUserNotConfirmedErr() *UserNotConfirmedErr {
	return &UserNotConfirmedErr{
		Message: "user not confirmed",
	}
}

func (e *UserNotConfirmedErr) Error() string {
	return e.Message
}

type InvalidCredentialsErr struct {
	Message string `json:"message"`
}

func NewInvalidCredentialsErr() *InvalidCredentialsErr {
	return &InvalidCredentialsErr{
		Message: "invalid credentials",
	}
}

func (e *InvalidCredentialsErr) Error() string {
	return e.Message
}

type InvalidCodeErr struct {
	Message string `json:"message"`
}

func NewInvalidCodeErr() *InvalidCodeErr {
	return &InvalidCodeErr{
		Message: "invalid code",
	}
}

func (e *InvalidCodeErr) Error() string {
	return e.Message
}

type ExpiredCodeErr struct {
	Message string `json:"message"`
}

func NewExpiredCodeErr() *ExpiredCodeErr {
	return &ExpiredCodeErr{
		Message: "code expired",
	}
}

func (e *ExpiredCodeErr) Error() string {
	return e.Message
}
