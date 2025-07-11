package message_type_common_model

type ButtonSubType string

const (
	QuickReply   ButtonSubType = "quick_reply"
	CallToAction ButtonSubType = "call_to_action"
)

// IsValid checks if the ButtonSubType is one of the allowed values.
func (bst ButtonSubType) IsValid() bool {
	switch bst {
	case QuickReply, CallToAction:
		return true
	default:
		return false
	}
}
