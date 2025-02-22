package client

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/charmingruby/impr/service/identity/internal/account/core/gateway"
	"github.com/charmingruby/impr/service/identity/pkg/awsc"
	"github.com/charmingruby/impr/service/identity/pkg/helper"
	"github.com/charmingruby/impr/service/identity/pkg/integration"
)

type CognitoIdentityProvider struct {
	cognito *awsc.CognitoClient
}

func NewCognitoIdentityProvider(cognito *awsc.CognitoClient) *CognitoIdentityProvider {
	return &CognitoIdentityProvider{
		cognito: cognito,
	}
}

func (c *CognitoIdentityProvider) SignUp(in gateway.SignUpInput) (*cognito.SignUpOutput, error) {
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
		return nil, err
	}

	return op, nil
}

func (c *CognitoIdentityProvider) ConfirmAccount(in gateway.ConfirmAccountInput) error {
	ctx, cancel := integration.NewContext()
	defer cancel()

	c.cognito.Client.ConfirmSignUp(ctx, &cognito.ConfirmSignUpInput{})

	return nil
}
