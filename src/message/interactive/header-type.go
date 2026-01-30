package interactive

type HeaderType string

const (
	Document HeaderType = "document"
	Image    HeaderType = "image"
	Video    HeaderType = "video"
	Text     HeaderType = "text"
)

// IsValid checks if the HeaderType is one of the allowed values.
func (ht HeaderType) IsValid() bool {
	switch ht {
	case Document, Image, Video, Text:
		return true
	default:
		return false
	}
}
