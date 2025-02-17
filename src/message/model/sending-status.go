package message_model

type SendingStatus string

const (
	Delivered SendingStatus = "delivered"
	Read      SendingStatus = "read"
	Sent      SendingStatus = "sent"

	Failed SendingStatus = "failed"
)
