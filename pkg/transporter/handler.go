package transporter

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h Handler) Handle(ctx context.Context, ev *events.CloudWatchEvent) error {
	fmt.Print(ev.Detail)
	return nil
}
