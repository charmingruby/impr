package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoClient struct {
	appClientID string
	client      *cognito.Client
}

func NewCognitoClient(appClientID string) (*CognitoClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	config, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := cognito.NewFromConfig(config)

	return &CognitoClient{
		appClientID: appClientID,
		client:      client,
	}, nil
}
