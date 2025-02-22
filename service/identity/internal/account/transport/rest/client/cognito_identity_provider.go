package client

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/charmingruby/impr/lib/pkg/awsc"
	"github.com/charmingruby/impr/lib/pkg/integration"
	"github.com/charmingruby/impr/service/identity/internal/account/core/gateway"
	"github.com/charmingruby/impr/service/identity/internal/account/core/model"
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

	parsedBirthdate := helper.BirthdateParser(in.Birthdate)

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
		return "", err
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
		return gateway.SignInOutput{}, err
	}

	return gateway.SignInOutput{
		AccessToken:  *op.AuthenticationResult.AccessToken,
		RefreshToken: *op.AuthenticationResult.RefreshToken,
	}, nil
}

func (c *CognitoIdentityProvider) RetrieveUser(accessToken string) (model.User, error) {
	ctx, cancel := integration.NewContext()
	defer cancel()

	op, err := c.cognito.Client.GetUser(ctx, &cognito.GetUserInput{
		AccessToken: &accessToken,
	})
	if err != nil {
		return model.User{}, err
	}

	u := cognitoUserToModel(op)

	return u, nil
}

func cognitoUserToModel(in *cognito.GetUserOutput) model.User {
	u := model.User{}

	for _, v := range in.UserAttributes {
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
			birthdate, err := time.Parse("2006-01-02", *v.Value)
			if err == nil {
				u.Birthdate = birthdate
			}
		}
	}

	return u
}
