package gateway

import "time"

type SignUpInput struct {
	FirstName string
	LastName  string
	Email     string
	Birthdate time.Time
	Password  string
}

type ConfirmAccountInput struct {
	Email string
	Code  string
}

type IdentityProvider interface {
	SignUp(in SignUpInput) error
	ConfirmAccount(in ConfirmAccountInput) error
	// SignIn()
	// RefreshSession()
	// ForgotPassword()
	// ResetPassword()
}
