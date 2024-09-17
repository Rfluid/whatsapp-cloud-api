package message_content_type_model

import message_type_interactive_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model/content-type/interactive"

// Described at
//
// https://developers.facebook.com/docs/whatsapp/on-premises/reference/messages#interactive-object
type Interactive struct {
	Type message_type_interactive_model.InteractiveType `json:"type"` // The type of interactive message you want to send.

	Header *message_type_interactive_model.Header `json:"header,omitempty"` // Header content displayed on top of a message. You cannot set a header if your interactive object is of product type.
	Body   *message_type_interactive_model.Body   `json:"body,omitempty"`   // Optional for type product. Required for other message types.
	Footer *message_type_interactive_model.Footer `json:"footer,omitempty"` // An object with the footer of the message.
	Action *message_type_interactive_model.Action `json:"action,omitempty"` // An action object with what you want the user to perform after reading the message. See action object for full information.
}

// Described at
//
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type ReceivedInteractive struct {
	Type        message_type_interactive_model.ReceivedInteractive `json:"type"` // The type of interactive message received.
	ListReply   *message_type_interactive_model.ListReplyData      `json:"list_reply,omitempty"`
	ButtonReply *message_type_interactive_model.ButtonReplyData    `json:"button_reply,omitempty"`
}
