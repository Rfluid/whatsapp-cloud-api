package webhook

import (
	"github.com/Rfluid/whatsapp-cloud-api/src/common"
	"github.com/Rfluid/whatsapp-cloud-api/src/message"
)

type Status struct {
	BizOpaqueCallbackData string                       `json:"biz_opaque_callback_data,omitempty"` // Arbitrary string used for tracking messages, groups of messages, you name it...
	Conversation          *Conversation                `json:"conversation,omitempty"`             // Information about the conversation.
	Errors                *[]common.Error        `json:"errors,omitempty"`                   // An array of error objects describing the error. Error objects have the following properties, which map to their equivalent properties in API error response payloads.
	ID                    string                       `json:"id"`                                 // The ID for the message that the business that is subscribed to the webhooks sent to a customer
	Pricing               *common.Pricing        `json:"pricing,omitempty"`                  // An object containing pricing information.
	RecipientID           string                       `json:"recipient_id"`                       // The customer's WhatsApp ID. A business can respond to a customer using this ID. This ID may not match the customer's phone number, which is returned by the API as input when sending a message to the customer.
	Status                *message.SendingStatus `json:"status,omitempty" validate:"omitempty,sending_status"`
	Timestamp             string                       `json:"timestamp"` // Date for the status message
}
