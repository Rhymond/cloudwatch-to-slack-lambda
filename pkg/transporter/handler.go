package transporter

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/slack-go/slack"
)

type Handler struct {
	client  *slack.Client
	channel string
}

func New(token, channel string) *Handler {
	h := new(Handler)
	h.client = slack.New(token)
	h.channel = channel
	return h
}

func (h Handler) Handle(ctx context.Context, ev *events.CloudwatchLogsEvent) error {
	events, err := ev.AWSLogs.Parse()
	if err != nil {
		return fmt.Errorf("unable to parse cloudwatch logs: %w", err)
	}

	for _, e := range events.LogEvents {
		var fields map[string]interface{}
		if err := json.Unmarshal([]byte(e.Message), &fields); err != nil {
			log.Printf("failed to unmarshal log message: %s", err)
			continue
		}

		attachements := make([]slack.AttachmentField, len(fields))
		for name, val := range fields {
			attachements = append(attachements, slack.AttachmentField{
				Title: name,
				Value: fmt.Sprint(val),
			})
		}

		attachment := slack.Attachment{
			Text:   "Error occured",
			Color:  "#f44336",
			Fields: attachements,
		}

		if _, _, err := h.client.PostMessage(
			h.channel,
			slack.MsgOptionAttachments(attachment),
			slack.MsgOptionAsUser(true),
		); err != nil {
			log.Printf("unable to send slack message: %s", err)
		}
	}

	return nil
}
