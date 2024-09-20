package template_model

type ButtonType string

var (
	PhoneNumber ButtonType = "PHONE_NUMBER"
	Url         ButtonType = "URL"
	QuickReply  ButtonType = "QUICK_REPLY"
	CopyCode    ButtonType = "COPY_CODE"
	Flow        ButtonType = "FLOW"
)
