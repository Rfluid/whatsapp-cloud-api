package profile_model

import common_model "github.com/Rfluid/whatsapp/src/common/model"

type BusinessProfile struct {
	About             string   `json:"about"`
	Address           string   `json:"address"`
	Description       string   `json:"description"`
	Email             string   `json:"email"`
	ProfilePictureURL string   `json:"profile_picture_url"`
	Websites          []string `json:"websites"`
	Vertical          string   `json:"vertical"`
	common_model.MessagingProduct
}
