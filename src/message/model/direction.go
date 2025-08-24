package message_model

import message_type_common_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model/content-type/common"

// Not the content per se of message. Just specifications of
// what and to whom will be sent.
type Direction struct {
	To   string                         `json:"to"`                                    // Whatsapp ID of receiver.
	Type message_type_common_model.Type `json:"type" validate:"required,message_type"` // Type of message.
}

type ReceivedDirection struct {
	ID   string `json:"id"`   // The ID for the message that was received by the business. You could use messages endpoint to mark this specific message as read.
	From string `json:"from"` // The customer's WhatsApp ID. A business can respond to a customer using this ID. This ID may not match the customer's phone number, which is returned by the API as input when sending a message to the customer.
}
