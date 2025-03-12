package client

import (
	"errors"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"

	"github.com/charmingruby/impr/lib/pkg/core/core_err"
	"github.com/charmingruby/impr/lib/pkg/http/client/integration"
	"github.com/charmingruby/impr/lib/pkg/http/client/service/awsc"
	"github.com/charmingruby/impr/service/identity/internal/account/core/gateway"
	"github.com/charmingruby/impr/service/identity/internal/account/core/model"
	"github.com/charmingruby/impr/service/identity/internal/shared/custom_err"
	"github.com/charmingruby/impr/service/identity/pkg/helper"
)

type CognitoIdentityProvider struct {
	cognito *awsc.CognitoClient
}

func NewCognitoIdentityProvider(cognito *awsc.CognitoClient) *CognitoIdentityProvider {
	return &CognitoIdentityProvider{
		cognito: cognito,
	}
}

func (c *CognitoIdentityProvider) SignUp(in gateway.SignUpInput) (string, error) {
	ctx, cancel := integration.NewContext()
	defer cancel()

	parsedBirthdate := helper.BirthdateToString(in.Birthdate)

	op, err := c.cognito.Client.SignUp(ctx, &cognito.SignUpInput{
		ClientId: &c.cognito.AppClientID,
		Username: &in.Email,
		Password: &in.Password,
		UserAttributes: []types.AttributeType{
			{
				Name: aws.String("given_name"), Value: aws.String(in.FirstName),
			},
			{
				Name: aws.String("family_name"), Value: aws.String(in.LastName),
			},
			{
				Name: aws.String("birthdate"), Value: aws.String(parsedBirthdate),
			},
		},
	})
	if err != nil {
		var invalidPasswordError *types.InvalidPasswordException
		if errors.As(err, &invalidPasswordError) {
			return "", core_err.NewInvalidFieldFormatErr("password", err)
		}

		var emailAlreadyExistsError *types.UsernameExistsException
		if errors.As(err, &emailAlreadyExistsError) {
			return "", core_err.NewConflictErr("email")
		}

		return "", custom_err.NewClientUncaughtErr(err)
	}

	return *op.UserSub, nil
}

func (c *CognitoIdentityProvider) ConfirmAccount(in gateway.ConfirmAccountInput) error {
	ctx, cancel := integration.NewContext()
	defer cancel()

	_, err := c.cognito.Client.ConfirmSignUp(ctx, &cognito.ConfirmSignUpInput{
		ClientId:         &c.cognito.AppClientID,
		Username:         &in.Email,
		ConfirmationCode: &in.Code,
	})

	if err != nil {
		var expiredCodeErr *types.ExpiredCodeException
		if errors.As(err, &expiredCodeErr) {
			return custom_err.NewExpiredCodeErr()
		}

		var mismatchCodeErr *types.CodeMismatchException
		if errors.As(err, &mismatchCodeErr) {
			return custom_err.NewInvalidCodeErr()
		}

		var invalidEmailErr *types.UserNotFoundException
		if errors.As(err, &invalidEmailErr) {
			return custom_err.NewInvalidCredentialsErr()
		}

		return custom_err.NewClientUncaughtErr(err)
	}

	return err
}

func (c *CognitoIdentityProvider) SignIn(in gateway.SignInInput) (gateway.SignInOutput, error) {
	ctx, cancel := integration.NewContext()
	defer cancel()

	op, err := c.cognito.Client.InitiateAuth(ctx, &cognito.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		ClientId: &c.cognito.AppClientID,
		AuthParameters: map[string]string{
			"USERNAME": in.Email,
			"PASSWORD": in.Password,
		},
	})
	if err != nil {
		var userNotConfirmedErr *types.UserNotConfirmedException
		if errors.As(err, &userNotConfirmedErr) {
			return gateway.SignInOutput{}, custom_err.NewUserNotConfirmedErr()
		}

		var invalidPasswordErr *types.NotAuthorizedException
		if errors.As(err, &invalidPasswordErr) {
			return gateway.SignInOutput{}, custom_err.NewInvalidCredentialsErr()
		}

		var invalidEmailErr *types.UserNotFoundException
		if errors.As(err, &invalidEmailErr) {
			return gateway.SignInOutput{}, custom_err.NewInvalidCredentialsErr()
		}

		return gateway.SignInOutput{}, custom_err.NewClientUncaughtErr(err)
	}

	return gateway.SignInOutput{
		AccessToken:  *op.AuthenticationResult.AccessToken,
		RefreshToken: *op.AuthenticationResult.RefreshToken,
	}, nil
}

func (c *CognitoIdentityProvider) RefreshSession(refreshToken string) (string, error) {
	ctx, cancel := integration.NewContext()
	defer cancel()

	op, err := c.cognito.Client.InitiateAuth(ctx, &cognito.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeRefreshTokenAuth,
		ClientId: &c.cognito.AppClientID,
		AuthParameters: map[string]string{
			"REFRESH_TOKEN": refreshToken,
		},
	})
	if err != nil {
		var invalidTokenErr *types.NotAuthorizedException
		if errors.As(err, &invalidTokenErr) {
			return "", custom_err.NewInvalidTokenErr()
		}

		return "", err
	}

	return *op.AuthenticationResult.AccessToken, nil
}

func (c *CognitoIdentityProvider) ForgotPassword(email string) error {
	ctx, cancel := integration.NewContext()
	defer cancel()

	_, err := c.cognito.Client.ForgotPassword(ctx, &cognito.ForgotPasswordInput{
		ClientId: &c.cognito.AppClientID,
		Username: &email,
	})
	if err != nil {
		var invalidEmailErr *types.UserNotFoundException
		if errors.As(err, &invalidEmailErr) {
			return core_err.NewResourceNotFoundErr("user")
		}

		return err
	}

	return nil
}

func (c *CognitoIdentityProvider) ResetPassword(in gateway.ResetPasswordInput) error {
	ctx, cancel := integration.NewContext()
	defer cancel()

	_, err := c.cognito.Client.ConfirmForgotPassword(ctx, &cognito.ConfirmForgotPasswordInput{
		ClientId:         &c.cognito.AppClientID,
		ConfirmationCode: &in.ConfirmationCode,
		Password:         &in.NewPassword,
		Username:         &in.Email,
	})
	if err != nil {
		var userNotFound *types.UserNotFoundException
		if errors.As(err, &userNotFound) {
			return core_err.NewResourceNotFoundErr("user")
		}

		var userNotConfirmedErr *types.UserNotConfirmedException
		if errors.As(err, &userNotConfirmedErr) {
			return custom_err.NewUserNotConfirmedErr()
		}

		var expiredCodeErr *types.ExpiredCodeException
		if errors.As(err, &expiredCodeErr) {
			return custom_err.NewExpiredCodeErr()
		}

		var mismatchCodeErr *types.CodeMismatchException
		if errors.As(err, &mismatchCodeErr) {
			return custom_err.NewInvalidCodeErr()
		}

		return err
	}

	return nil
}

func (c *CognitoIdentityProvider) FindUserByID(id string) (model.User, error) {
	ctx, cancel := integration.NewContext()
	defer cancel()

	op, err := c.cognito.Client.ListUsers(ctx, &cognito.ListUsersInput{
		UserPoolId: &c.cognito.UserPoolID,
		Filter:     aws.String("sub = \"" + id + "\""),
	})
	if err != nil {
		return model.User{}, err
	}

	if len(op.Users) == 0 {
		return model.User{}, core_err.NewResourceNotFoundErr("user")
	}

	foundUser := op.Users[0]

	u := model.User{}

	for _, v := range foundUser.Attributes {
		switch *v.Name {
		case "sub":
			u.ID = *v.Value
		case "email":
			u.Email = *v.Value
		case "email_verified":
			isVerified, _ := strconv.ParseBool(*v.Value)
			u.IsVerified = isVerified
		case "given_name":
			u.FirstName = *v.Value
		case "family_name":
			u.LastName = *v.Value
		case "birthdate":
			birthdate, err := helper.StringToBirthdate(*v.Value)
			if err == nil {
				u.Birthdate = birthdate
			}
		}
	}

	u.CreatedAt = *foundUser.UserCreateDate
	u.UpdatedAt = foundUser.UserLastModifiedDate

	return u, nil
}

func (c *CognitoIdentityProvider) RetrieveUserAttributesFromToken(accessToken string) (gateway.RetrieveUserAttributesFromTokenOutput, error) {
	ctx, cancel := integration.NewContext()
	defer cancel()

	op, err := c.cognito.Client.GetUser(ctx, &cognito.GetUserInput{
		AccessToken: &accessToken,
	})
	if err != nil {
		var userNotFound *types.UserNotFoundException
		if errors.As(err, &userNotFound) {
			return gateway.RetrieveUserAttributesFromTokenOutput{}, core_err.NewResourceNotFoundErr("user")
		}

		var userNotConfirmedErr *types.UserNotConfirmedException
		if errors.As(err, &userNotConfirmedErr) {
			return gateway.RetrieveUserAttributesFromTokenOutput{}, custom_err.NewUserNotConfirmedErr()
		}

		return gateway.RetrieveUserAttributesFromTokenOutput{}, err
	}

	uAtt := gateway.RetrieveUserAttributesFromTokenOutput{}

	for _, v := range op.UserAttributes {
		switch *v.Name {
		case "sub":
			uAtt.ID = *v.Value
		case "email":
			uAtt.Email = *v.Value
		case "email_verified":
			isVerified, _ := strconv.ParseBool(*v.Value)
			uAtt.IsVerified = isVerified
		case "given_name":
			uAtt.FirstName = *v.Value
		case "family_name":
			uAtt.LastName = *v.Value
		case "birthdate":
			birthdate, err := time.Parse("2006-01-02", *v.Value)
			if err == nil {
				uAtt.Birthdate = birthdate
			}
		}
	}
	return uAtt, nil
}
