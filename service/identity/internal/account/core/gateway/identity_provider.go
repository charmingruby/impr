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

type SignInInput struct {
	Email    string
	Password string
}

type SignInOutput struct {
	AccessToken  string
	RefreshToken string
}

type IdentityProvider interface {
	SignUp(in SignUpInput) (string, error)
	ConfirmAccount(in ConfirmAccountInput) error
	SignIn(in SignInInput) (SignInOutput, error)
	// RefreshSession()
	// ForgotPassword()
	// ResetPassword()
}
