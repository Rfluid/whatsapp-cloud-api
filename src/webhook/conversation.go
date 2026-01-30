package webhook

import "github.com/Rfluid/whatsapp-cloud-api/src/common"

type Conversation struct {
	ID                  string `json:"id"`     // Represents the ID of the conversation the given status notification belongs to.
	Origin              Origin `json:"origin"` // Describes conversation category
	ExpirationTimestamp string `json:"expiration_timestamp"`
}

type Origin struct {
	Type common.PricingCategory `json:"type"` // Indicates conversation category. This can also be referred to as a conversation entry point
}
