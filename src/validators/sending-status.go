package validators

import (
	"strings"

	"github.com/Rfluid/whatsapp-cloud-api/src/message"
	"github.com/go-playground/validator/v10"
)

// sendingStatusValidation checks if a string is a valid SendingStatus (case-insensitive).
func sendingStatusValidation(fl validator.FieldLevel) bool {
	input := fl.Field().String()

	status := message.SendingStatus(strings.ToLower(input))
	return status.IsValid()
}

// RegisterSendingStatusValidator registers the "sending_status" validator.
func RegisterSendingStatusValidator(v *validator.Validate) error {
	return v.RegisterValidation("sending_status", sendingStatusValidation)
}
