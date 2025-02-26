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

type RetrieveUserAttributesFromTokenOutput struct {
	ID         string
	Email      string
	FirstName  string
	LastName   string
	IsVerified bool
	Birthdate  time.Time
}

type IdentityProvider interface {
	SignUp(in SignUpInput) (string, error)
	ConfirmAccount(in ConfirmAccountInput) error
	SignIn(in SignInInput) (SignInOutput, error)
	RefreshSession(refreshToken string) (string, error)
	FindUserByID(id string) (model.User, error)
	RetrieveUserAttributesFromToken(accessToken string) (RetrieveUserAttributesFromTokenOutput, error)

	ForgotPassword(email string) error
	ResetPassword(in ResetPasswordInput) error
}
