package message_type_common_model

type ButtonSubType string

const (
	BookingConfirmation ButtonSubType = "BOOKING_CONFIRMATION"
	Catalog             ButtonSubType = "CATALOG"
	CopyCode            ButtonSubType = "COPY_CODE"
	Flow                ButtonSubType = "FLOW"
	Mpm                 ButtonSubType = "MPM"
	OrderDetails        ButtonSubType = "ORDER_DETAILS"
	QuickReply          ButtonSubType = "QUICK_REPLY"
	Reminder            ButtonSubType = "REMINDER"
	Url                 ButtonSubType = "URL"
	VoiceCall           ButtonSubType = "VOICE_CALL"
)

// IsValid checks if the ButtonSubType is one of the allowed values.
func (bst ButtonSubType) IsValid() bool {
	switch bst {
	case BookingConfirmation,
		Catalog,
		CopyCode,
		Flow,
		Mpm,
		OrderDetails,
		QuickReply,
		Reminder,
		Url,
		VoiceCall:
		return true
	default:
		return false
	}
}
