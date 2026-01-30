// Provides models to handle number registering validation.
package phone

import "github.com/Rfluid/whatsapp-cloud-api/src/common"

// To register new business phone number.
type RegisterPayload struct {
	DataLocalizationRegion string `json:"data_localization_region,omitempty"` // If included, enables local storage on the business phone number. Value must be a 2-letter ISO 3166 country code (e.g. IN) indicating the country where you want data-at-rest to be stored.
	common.MessagingProduct
	Pin
}
