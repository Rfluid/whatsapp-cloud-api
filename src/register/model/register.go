// Provides models to handle number registering validation.
package register_model

import (
	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
	two_step_verification_model "github.com/Rfluid/whatsapp-cloud-api/src/two-step-verification/model"
)

// To register new business phone number.
type Register struct {
	DataLocalizationRegion string `json:"data_localization_region,omitempty"` // If included, enables local storage on the business phone number. Value must be a 2-letter ISO 3166 country code (e.g. IN) indicating the country where you want data-at-rest to be stored.
	common_model.MessagingProduct
	two_step_verification_model.Pin
}
