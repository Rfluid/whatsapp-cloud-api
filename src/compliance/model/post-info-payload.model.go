package compliance_model

type PostInfoPayload struct {
	MessagingProduct        string                  `json:"messaging_product"`
	EntityName              string                  `json:"entity_name"`
	EntityType              string                  `json:"entity_type"`
	IsRegistered            bool                    `json:"is_registered"`
	GrievanceOfficerDetails GrievanceOfficerDetails `json:"grievance_officer_details"`
	CustomerCareDetails     CustomerCareDetails     `json:"customer_care_details"`
}
