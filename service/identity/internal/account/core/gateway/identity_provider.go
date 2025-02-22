package gateway

import (
	"time"

	"github.com/charmingruby/impr/service/identity/internal/account/core/model"
)

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

type ResetPasswordInput struct {
	Email            string
	NewPassword      string
	ConfirmationCode string
}

type IdentityProvider interface {
	SignUp(in SignUpInput) (string, error)
	ConfirmAccount(in ConfirmAccountInput) error
	SignIn(in SignInInput) (SignInOutput, error)
	RefreshSession(refreshToken string) (string, error)
	ForgotPassword(email string) error
	ResetPassword(in ResetPasswordInput) error
	RetriveUser(accessToken string) (model.User, error)
}
