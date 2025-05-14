package webhook_model

import common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"

type Conversation struct {
	Id                  string `json:"id"`     // Represents the ID of the conversation the given status notification belongs to.
	Origin              Origin `json:"origin"` // Describes conversation category
	ExpirationTimestamp string `json:"expiration_timestamp"`
}

type Origin struct {
	Type common_model.PricingCategory `json:"type"` // Indicates conversation category. This can also be referred to as a conversation entry point
}
