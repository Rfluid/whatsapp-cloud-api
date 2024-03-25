package message_model

// Provided when you want to answer a message.
type Context struct {
	MessageId string `json:"message_id"` // Id of the message you want to answer.
}
