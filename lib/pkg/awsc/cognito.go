package awsc

import (
	"github.com/aws/aws-sdk-go-v2/config"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/charmingruby/impr/lib/pkg/integration"
)

type CognitoClient struct {
	AppClientID string
	Client      *cognito.Client
}

func NewCognitoClient(appClientID string) (*CognitoClient, error) {
	ctx, cancel := integration.NewContext()
	defer cancel()

	config, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := cognito.NewFromConfig(config)

	return &CognitoClient{
		AppClientID: appClientID,
		Client:      client,
	}, nil
}
