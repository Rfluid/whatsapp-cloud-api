// Modules to handle WhatsApp business profile validation.
package profile_model

type BusinessProfileField string

const (
	About             BusinessProfileField = "about"
	Address           BusinessProfileField = "address"
	Description       BusinessProfileField = "description"
	Email             BusinessProfileField = "email"
	MessagingProduct  BusinessProfileField = "messaging_product"
	ProfilePictureURL BusinessProfileField = "profile_picture_url"
	Websites          BusinessProfileField = "websites"
	Vertical          BusinessProfileField = "vertical"
)
