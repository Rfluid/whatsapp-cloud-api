package message_model

// Contact data returned in response.
type ResponseContact struct {
	Input string `json:"input"` // User's phone number.
	WAId  string `json:"wa_id"` // User's WhatsApp id returned.
}
