package message_model

// Not the content per se of message. Just specifications of
// what and to whom will be sent.
type Direction struct {
	To   string `json:"to"`   // Whatsapp ID of receiver.
	Type Type   `json:"type"` // Type of message.
}
