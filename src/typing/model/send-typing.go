package typing_model

import (
	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
	message_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model"
)

type SendTyping struct {
	MessageID       string                      `json:"message_id" validate:"required"`
	Status          message_model.SendingStatus `json:"status" validate:"required,sending_status"`
	TypingIndicator TypingIndicator             `json:"typing_indicator" validate:"required"`

	common_model.MessagingProduct
}

type TypingIndicatorType string

var Text TypingIndicatorType = "text"

// IsValid checks if the TypingIndicatorType is one of the allowed values.
func (s TypingIndicatorType) IsValid() bool {
	switch s {
	case Text:
		return true
	default:
		return false
	}
}

type TypingIndicator struct {
	Type TypingIndicatorType `json:"type" validate:"required,typing_indicator_type"`
}
