package message_model

import common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"

// Entire message that can be sent.
//
// Useful guide at https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages.
type Message struct {
	RecipientIdentityKeyHash string   `json:"recipient_identity_key_hash,omitempty"` // To use it you need to allow identity check at code.service. If the field is provided, this message will only be sent if "recipient_identity_key_hash" matches client's current hash.
	BizOpaqueCallbackData    string   `json:"biz_opaque_callback_data,omitempty"`    // Arbitrary string used for tracking messages, groups of messages, you name it...
	Context                  *Context `json:"context,omitempty"`                     // Used to answer a message.
	Default
	Direction
	Content
}

// Entire message that can be receive.
//
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/components#messages-object
type MessageReceived struct {
	Context   *ReceivedContext      `json:"context,omitempty"`
	Errors    *[]common_model.Error `json:"errors,omitempty"`
	Referral  *Referral             `json:"referral,omitempty"`
	System    *System               `json:"system,omitempty"`
	Timestamp string                `json:"timestamp"` // Unix timestamp indicating when the WhatsApp server received the message from the customer.
	Type      ReceiveType           `json:"type"`
	Content
	ReceivedDirection
}
