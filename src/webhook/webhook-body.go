package webhook

import (
	"github.com/Rfluid/whatsapp-cloud-api/src/common"
	"github.com/Rfluid/whatsapp-cloud-api/src/message"
)

// Body of the received webhook.
//
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/components
type WebhookBody struct {
	Object common.BusinessAccount `json:"object"` // The specific webhook a business is subscribed to. The webhook is whatsapp_business_account.
	Entry  []Entry                      `json:"entry"`
}

type Entry struct {
	ID      string   `json:"id"`
	Changes []Change `json:"changes"`
}

type Change struct {
	Value Value `json:"value"`
	Field Field `json:"field"`
}

type Value struct {
	Metadata Metadata `json:"metadata"`

	common.MessagingProduct

	// Specific webhooks payload.

	Contacts *[]Contact                       `json:"contacts,omitempty"`
	Errors   *[]common.Error            `json:"errors,omitempty"`
	Messages *[]message.MessageReceived `json:"messages,omitempty"`
	Statuses *[]Status                        `json:"statuses,omitempty"`
}

type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}
