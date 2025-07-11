package validators

import (
	"strings"

	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
	"github.com/go-playground/validator/v10"
)

// supportedMimeTypeValidation checks if the given string is a supported MIME type.
func supportedMimeTypeValidation(fl validator.FieldLevel) bool {
	input, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	mimeType := common_model.SupportedMimeTypes(strings.ToLower(input))
	return mimeType.IsValid()
}

// RegisterSupportedMimeTypeValidator registers the "supported_mime_type" validator.
func RegisterSupportedMimeTypeValidator(v *validator.Validate) error {
	return v.RegisterValidation("supported_mime_type", supportedMimeTypeValidation)
}
