package message_type_common_model

type Button struct {
	Payload string `json:"payload,omitempty"`
	Text    string `json:"text,omitempty"`
}
