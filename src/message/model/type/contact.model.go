package message_type_model

import (
	"time"

	message_type_contact_model "github.com/Rfluid/whatsapp/src/message/model/type/contact"
)

type Contact struct {
	Addresses []message_type_contact_model.Address `json:"addresses,omitempty"`
	Birthday  time.Time                            `json:"birthday,omitempty"`
	Emails    []message_type_contact_model.Email   `json:"emails,omitempty"`
	Name      message_type_contact_model.Name      `json:"name,omitempty"`
	Org       message_type_contact_model.Org       `json:"org,omitempty"`
	Phones    []message_type_contact_model.Phone   `json:"phones,omitempty"`
	URLs      []message_type_contact_model.URL     `json:"urls,omitempty"`
}
