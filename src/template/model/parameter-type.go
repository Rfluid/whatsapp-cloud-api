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

// IsValid checks if the ParameterType is one of the allowed types.
func (pt ParameterType) IsValid() bool {
	switch pt {
	case Text, Currency, DateTime, Image, Video, Sticker, Document, Button, Payload:
		return true
	default:
		return false
	}
}
