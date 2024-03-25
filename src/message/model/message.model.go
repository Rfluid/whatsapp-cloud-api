package message_model

// Entire message that can be sent.
//
// Useful guide at https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages.
type Message struct {
	RecipientIdentityKeyHash string   `json:"recipient_identity_key_hash,omitempty"` // To use it you need to allow identity check at code.service. If the field is provided, this message will only be sent if "recipient_identity_key_hash" matches client's current hash.
	BizOpaqueCallbackData    string   `json:"biz_opaque_callback_data,omitempty"`    // Random string used for tracking messages, groups of messages, you name it...
	Context                  *Context `json:"context,omitempty"`                     // Used to answer a message.
	Default
	Direction
	Content
}
