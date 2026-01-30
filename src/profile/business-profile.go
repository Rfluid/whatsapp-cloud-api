package profile

import "github.com/Rfluid/whatsapp-cloud-api/src/common"

type BusinessProfile struct {
	About             string   `json:"about"`
	Address           string   `json:"address"`
	Description       string   `json:"description"`
	Email             string   `json:"email"`
	ProfilePictureURL string   `json:"profile_picture_url"`
	Websites          []string `json:"websites"`
	Vertical          string   `json:"vertical"`
	common.MessagingProduct
}
