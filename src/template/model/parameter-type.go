package template_model

type ParameterType string

const (
	Text     ParameterType = "text"
	Currency ParameterType = "currency"
	DateTime ParameterType = "date_time"
	Image    ParameterType = "image"
	Video    ParameterType = "video"
	Sticker  ParameterType = "sticker"
	Document ParameterType = "document"
	Button   ParameterType = "button"
	Payload  ParameterType = "payload"
)
