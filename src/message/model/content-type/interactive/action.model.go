package message_type_interactive_model

// Described at
//
// https://developers.facebook.com/docs/whatsapp/on-premises/reference/messages#action-object
type Action struct {
	Button   string        `json:"button,omitempty"`   // Required for list messages. It cannot be an empty string and must be unique within the message. Emojis are supported, markdown is not. Maximum length: 20 characters.
	Buttons  *[]ButtonData `json:"buttons,omitempty"`  // Required for list messages. Maximum number of buttons is 3.
	Sections *[]Section    `json:"sections,omitempty"` // There is a minimum of 1 and maximum of 10.
}
