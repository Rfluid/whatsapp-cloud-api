package message_type_interactive_model

import message_type_common_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model/content-type/common"

// Described at
//
// https://developers.facebook.com/docs/whatsapp/on-premises/reference/messages#body-object
type Body struct {
	Text message_type_common_model.TextData `json:"text"` // Required if body is present
}
