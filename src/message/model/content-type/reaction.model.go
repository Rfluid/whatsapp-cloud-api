package message_content_type_model

import (
	"errors"
	"strings"
)

// Reaction messages.
type Reaction struct {
	MessageId string `json:"message_id"` // Id of the message to react to. Something like "wamid.HBgLM..."
	Emoji     string `json:"emoji"`      // The emoji encoding. Something like "\uD83D\uDE00"
}

// Validate wamid format.
func (r *Reaction) Validate() error {
	messageSplitted := strings.Split(r.MessageId, ".")

	ch := make(chan error)

	go func() {
		if messageSplitted[0] != "wamid" {
			ch <- errors.New("the message_id of your reaction is not valid. It should start with \"wamid\"")
		} else {
			ch <- nil
		}
	}()

	go func() {
		if len(messageSplitted) < 2 {
			ch <- errors.New("the message_id of your reaction is not valid. It should be something like \"wamid.HBgLM...\"")
		} else {
			ch <- nil
		}
	}()

	if err := <-ch; err != nil {
		return err
	}

	return nil
}
