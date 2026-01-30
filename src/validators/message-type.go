package validators

import (
	"strings"

	"github.com/Rfluid/whatsapp-cloud-api/src/message/content"
	"github.com/go-playground/validator/v10"
)

// messageTypeValidation validates if a string maps to a valid message Type (case-insensitive).
func messageTypeValidation(fl validator.FieldLevel) bool {
	input := fl.Field().String()

	msgType := content.Type(strings.ToLower(input))
	return msgType.IsValid()
}

// RegisterMessageTypeValidator registers the custom "message_type" validator.
func RegisterMessageTypeValidator(v *validator.Validate) error {
	return v.RegisterValidation("message_type", messageTypeValidation)
}
