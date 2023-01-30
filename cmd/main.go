package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rhymond/cloudwatchtoslack/pkg/transporter"
)

func main() {
	ctx := context.Background()
	t := transporter.New(
		os.Getenv("SLACK_TOKEN"),
		os.Getenv("CHANNEL_ID"),
	)
	lambda.StartWithContext(ctx, t.Handle)
}
