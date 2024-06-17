package message_model

// Provided when you want to answer a message.
type Context struct {
	MessageId string `json:"message_id"` // Id of the message you want to answer.
}

// Only included when a user replies or interacts with one of your messages. Context objects can have the following properties:
type ReceivedContext struct {
	Forwarded           bool            `json:"forwarded"`            // Set to true if the message received by the business has been forwarded.
	FrequentlyForwarded bool            `json:"frequently_forwarded"` // Set to true if the message received by the business has been forwarded more than 5 times.
	From                string          `json:"from"`                 // The WhatsApp ID for the customer who replied to an inbound message.
	Id                  string          `json:"id"`                   // The message ID for the sent message for an inbound reply.
	ReferredProduct     ReferredProduct `json:"referred_product"`     // Referred product object describing the product the user is requesting information about. You must parse this value if you support Product Enquiry Messages.
}
