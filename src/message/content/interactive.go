package content

import "github.com/Rfluid/whatsapp-cloud-api/src/message/interactive"

// Described at
//
// https://developers.facebook.com/docs/whatsapp/on-premises/reference/messages#interactive-object
type Interactive struct {
	Type interactive.InteractiveType `json:"type"` // The type of interactive message you want to send.

	Header *interactive.Header `json:"header,omitempty"` // Header content displayed on top of a message. You cannot set a header if your interactive object is of product type.
	Body   *interactive.Body   `json:"body,omitempty"`   // Optional for type product. Required for other message types.
	Footer *interactive.Footer `json:"footer,omitempty"` // An object with the footer of the message.
	Action *interactive.Action `json:"action,omitempty"` // An action object with what you want the user to perform after reading the message. See action object for full information.
}

// Described at
//
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type ReceivedInteractive struct {
	Type        interactive.ReceivedInteractive `json:"type"` // The type of interactive message received.
	ListReply   *interactive.ListReplyData      `json:"list_reply,omitempty"`
	ButtonReply *interactive.ButtonReplyData    `json:"button_reply,omitempty"`
}
