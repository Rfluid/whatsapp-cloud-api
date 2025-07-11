package template_model

type ButtonType string

const (
	PhoneNumber ButtonType = "PHONE_NUMBER"
	Url         ButtonType = "URL"
	QuickReply  ButtonType = "QUICK_REPLY"
	CopyCode    ButtonType = "COPY_CODE"
	Flow        ButtonType = "FLOW"
)

// IsValid returns true if the ButtonType is one of the predefined constants.
func (bt ButtonType) IsValid() bool {
	switch bt {
	case PhoneNumber, Url, QuickReply, CopyCode, Flow:
		return true
	default:
		return false
	}
}
