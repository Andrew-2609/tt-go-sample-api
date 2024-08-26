package messaging

import (
	"bytes"
	"context"
	"encoding/json"
	"tt-go-sample-api/external/aws/sqs"
	"tt-go-sample-api/pkg/logger"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsSqs "github.com/aws/aws-sdk-go-v2/service/sqs"
)

// AWSSQSProducer is an AWS SQS message producer.
// It implements the messaging.MessageProducer
// interface.
type AWSSQSProducer struct {
	queueUrl string
}

// NewAWSSQSProducer returns a pointer of
// AWSSQSProducer, with the given Queue URL.
func NewAWSSQSProducer(queueUrl string) *AWSSQSProducer {
	return &AWSSQSProducer{queueUrl: queueUrl}
}

// Produce produces an SQS message.
//
// Currently, it only supports FIFO messages.
func (p *AWSSQSProducer) Produce(ctx context.Context, message Message) error {
	messageBytes := new(bytes.Buffer)

	if err := json.NewEncoder(messageBytes).Encode(message.Payload); err != nil {
		return err
	}

	_, err := sqs.GetAPISQSSingletonSingleton().Client().SendMessage(ctx, &awsSqs.SendMessageInput{
		QueueUrl:               &p.queueUrl,
		MessageBody:            aws.String(messageBytes.String()),
		MessageGroupId:         aws.String(message.GroupID),
		MessageDeduplicationId: aws.String(message.ID),
	})

	if err != nil {
		return err
	}

	logger.APILoggerSingleton.Info(ctx, logger.LogInput{
		Message: "Message successfuly sent to SQS",
	})

	return nil
}
