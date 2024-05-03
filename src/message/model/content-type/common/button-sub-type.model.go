package message_type_common_model

type ButtonSubType string

const (
	QuickReply   ButtonSubType = "quick_reply"
	CallToAction ButtonSubType = "call_to_action"
)
