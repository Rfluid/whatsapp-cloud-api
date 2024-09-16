package message_content_type_model

import message_type_interactive_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model/content-type/interactive"

// Described at
//
// https://developers.facebook.com/docs/whatsapp/on-premises/reference/messages#interactive-object
type Interactive struct {
	Header *message_type_interactive_model.Header `json:"header,omitempty"`
	Body   *message_type_interactive_model.Body   `json:"body,omitempty"`
	Footer *message_type_interactive_model.Footer `json:"footer,omitempty"`
	Action *message_type_interactive_model.Action `json:"action,omitempty"`
}
