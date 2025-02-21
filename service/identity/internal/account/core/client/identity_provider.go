package client

import "time"

type SignUpInput struct {
	FirstName string
	LastName  string
	Email     string
	Birthdate time.Time
	Password  string
}

type IdentityProvider interface {
	SignUp(in SignUpInput) error
	// ConfirmAccount()
	// SignIn()
	// RefreshSession()
	// ForgotPassword()
	// ResetPassword()
}
