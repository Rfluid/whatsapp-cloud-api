package interactive

import (
	"github.com/Rfluid/whatsapp-cloud-api/src/media"
)

// Described at
//
// https://developers.facebook.com/docs/whatsapp/on-premises/reference/messages#header-object
type Header struct {
	Type HeaderType `json:"type" validate:"required,interactive_header_type"` // Type of message.

	Text     *string               `json:"text,omitempty"`
	Image    *media.UseMedia `json:"image,omitempty"`
	Video    *media.UseMedia `json:"video,omitempty"`
	Document *media.UseMedia `json:"document,omitempty"`
	Audio    *media.UseMedia `json:"audio,omitempty"`
}
