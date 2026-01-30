// Provides models to handle messages.
package message

import (
	"github.com/Rfluid/whatsapp-cloud-api/src/media"
	"github.com/Rfluid/whatsapp-cloud-api/src/message/content"
	"github.com/Rfluid/whatsapp-cloud-api/src/template"
)

// The message real content seen by user.
type Content struct {
	Text        *content.TextData     `json:"text,omitempty"`
	Reaction    *content.ReactionData `json:"reaction,omitempty"`
	Image       *media.UseMedia                   `json:"image,omitempty"`
	Video       *media.UseMedia                   `json:"video,omitempty"`
	Document    *media.UseMedia                   `json:"document,omitempty"`
	Audio       *media.UseMedia                   `json:"audio,omitempty"`
	Sticker     *media.UseMedia                   `json:"sticker,omitempty"`
	Location    *content.LocationData `json:"location,omitempty"`
	Template    *template.UseTemplate             `json:"template,omitempty"`
	Interactive *content.Interactive `json:"interactive,omitempty"`
	Contacts    *[]content.Contact   `json:"contacts,omitempty"`
	Button      *content.ButtonData   `json:"button,omitempty"`
	Order       *content.OrderData    `json:"order,omitempty"`
}

type ReceivedContent struct {
	Text        *content.TextData             `json:"text,omitempty"`
	Reaction    *content.ReactionData         `json:"reaction,omitempty"`
	Image       *media.UseMedia                           `json:"image,omitempty"`
	Video       *media.UseMedia                           `json:"video,omitempty"`
	Document    *media.UseMedia                           `json:"document,omitempty"`
	Audio       *media.UseMedia                           `json:"audio,omitempty"`
	Sticker     *media.UseMedia                           `json:"sticker,omitempty"`
	Location    *content.LocationData         `json:"location,omitempty"`
	Template    *template.UseTemplate                     `json:"template,omitempty"`
	Interactive *content.ReceivedInteractive `json:"interactive,omitempty"`
	Contacts    *[]content.Contact           `json:"contacts,omitempty"`
	Button      *content.ButtonData           `json:"button,omitempty"`
	Order       *content.OrderData            `json:"order,omitempty"`
}
