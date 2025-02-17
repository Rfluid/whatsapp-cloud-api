package media_model

// Used to retrieve media info.
type RetrieveInfo struct {
	PhoneNumberId string `json:"phone_number_id,omitempty"` // Business phone number ID. If included, the operation will only be processed if the ID matches the ID of the business phone number that the media was uploaded on.
}
