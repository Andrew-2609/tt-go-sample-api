package sqs

import (
	"context"
	"tt-go-sample-api/config"
	"tt-go-sample-api/util"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	awsSqs "github.com/aws/aws-sdk-go-v2/service/sqs"
)

// APISQS is a wrapper of an AWS SQS Client.
// It shall be used to perform SQS operations
// throughout the application.
type APISQS struct {
	client *awsSqs.Client
}

// NewAPISQS returns a pointer to APISQS.
func NewAPISQS() *APISQS {
	return &APISQS{}
}

// Client returns the APISQS's inner AWS SQS
// client, that can be used to run SQS
// commands.
func (d *APISQS) Client() *awsSqs.Client {
	return d.client
}

// Connect tries to connect to the APISQS's
// inner SQS client, loading its configuration
// based on the API's configuration.
func (a *APISQS) Connect(ctx context.Context, apiConfig *config.APIConfig) error {
	if a.client != nil {
		return nil
	}

	sqsConfig, err := getSQSConfig(ctx, apiConfig)

	if err != nil {
		return err
	}

	a.client = awsSqs.NewFromConfig(sqsConfig)

	return nil
}

// getSQSConfig formats the needed SQS
// configuration in order to have a healthy
// client that the API can use.
func getSQSConfig(ctx context.Context, apiConfig *config.APIConfig) (aws.Config, error) {
	sqsConfig, err := awsConfig.LoadDefaultConfig(ctx, awsConfig.WithRegion(apiConfig.AWSRegion))

	if !util.IsProductionEnv() {
		sqsConfig.BaseEndpoint = aws.String(apiConfig.AWSEndpoint)
		sqsConfig.Credentials = credentials.NewStaticCredentialsProvider("test", "test", "")
	}

	if err != nil {
		return aws.Config{}, err
	}

	return sqsConfig, nil
}
