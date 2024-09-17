package message_type_interactive_model

// Described at
//
// https://developers.facebook.com/docs/whatsapp/on-premises/reference/messages#button-object
type ButtonData struct {
	Type  ButtonType      `json:"type"`  // The only supported option is reply for Reply Button Messages.
	Reply ButtonReplyData `json:"reply"` // The reply object contains the title and ID of the reply button.
}

type ButtonReplyData struct {
	Title string `json:"title"` //  It cannot be an empty string and must be unique within the message. Emojis are supported, markdown is not. Maximum length: 20 characters.
	Id    string `json:"id"`    // Unique identifier for your button. This ID is returned in the webhook when the button is clicked by the user. Maximum length: 256 characters.
}
