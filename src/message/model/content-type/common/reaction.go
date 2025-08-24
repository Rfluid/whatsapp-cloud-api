package message_type_common_model

import (
	"errors"
	"strings"
	"sync"
)

// Reaction messages.
type ReactionData struct {
	MessageID string `json:"message_id"` // ID of the message to react to. Something like "wamid.HBgLM..."
	Emoji     string `json:"emoji"`      // The emoji encoding. Something like "\uD83D\uDE00"
}

// Validate wamid format.
func (r *ReactionData) Validate() error {
	messageSplitted := strings.Split(r.MessageID, ".")

	ch := make(chan error, 2)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if messageSplitted[0] != "wamid" {
			ch <- errors.New("the message_id of your reaction is not valid. It should start with \"wamid\"")
		} else {
			ch <- nil
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if len(messageSplitted) < 2 {
			ch <- errors.New("the message_id of your reaction is not valid. It should be something like \"wamid.HBgLM...\"")
		} else {
			ch <- nil
		}
	}()

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	for err := range ch {
		if err != nil {
			return err
		}
	}

	return nil
}
