package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rhymond/cloudwatchtoslack/pkg/transporter"
)

func main() {
	ctx := context.Background()
	t := transporter.New()
	lambda.StartWithContext(ctx, t.Handle)
}
