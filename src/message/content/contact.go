// Handles content validation.
package content

import (
	"time"

	"github.com/Rfluid/whatsapp-cloud-api/src/message/contact"
)

type Contact struct {
	Addresses []contact.Address `json:"addresses,omitempty"`
	Birthday  *time.Time                           `json:"birthday,omitempty"`
	Emails    []contact.Email   `json:"emails,omitempty"`
	Name      contact.Name      `json:"name,omitempty"`
	Org       contact.Org       `json:"org,omitempty"`
	Phones    []contact.Phone   `json:"phones,omitempty"`
	URLs      []contact.URL     `json:"urls,omitempty"`
}
