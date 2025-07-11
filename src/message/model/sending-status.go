package message_model

type SendingStatus string

const (
	Delivered SendingStatus = "delivered"
	Read      SendingStatus = "read"
	Sent      SendingStatus = "sent"

	Failed SendingStatus = "failed"
)

// IsValid checks if the SendingStatus is one of the allowed values.
func (s SendingStatus) IsValid() bool {
	switch s {
	case Delivered, Read, Sent, Failed:
		return true
	default:
		return false
	}
}
