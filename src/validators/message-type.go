package validators

import (
	"strings"

	message_type_common_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model/content-type/common"
	"github.com/go-playground/validator/v10"
)

// messageTypeValidation validates if a string maps to a valid message Type (case-insensitive).
func messageTypeValidation(fl validator.FieldLevel) bool {
	input := fl.Field().String()

	msgType := message_type_common_model.Type(strings.ToLower(input))
	return msgType.IsValid()
}

// RegisterMessageTypeValidator registers the custom "message_type" validator.
func RegisterMessageTypeValidator(v *validator.Validate) error {
	return v.RegisterValidation("message_type", messageTypeValidation)
}
