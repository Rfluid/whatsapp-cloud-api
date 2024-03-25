package message_model

// Entire message that can be sent.
type Message struct {
	RecipientIdentityKeyHash string `json:"recipient_identity_key_hash,omitempty"` // To use it you need to allow identity check at code.service. If the field is provided, this message will only be sent if "recipient_identity_key_hash" matches client's current hash.
	Default
	Way
	Content
}
