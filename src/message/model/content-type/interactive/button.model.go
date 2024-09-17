package message_type_interactive_model

// Described at
//
// https://developers.facebook.com/docs/whatsapp/on-premises/reference/messages#button-object
type ButtonData struct {
	Type  ButtonType      `json:"type"`  // The only supported option is reply for Reply Button Messages.
	Reply ButtonReplyData `json:"reply"` // The reply object contains the title and ID of the reply button.
}
