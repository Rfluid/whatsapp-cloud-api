package compliance_model

import common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"

type InfoData struct {
	EntityName              string                  `json:"entity_name"`
	EntityType              string                  `json:"entity_type"`
	IsRegistered            bool                    `json:"is_registered"`
	GrievanceOfficerDetails GrievanceOfficerDetails `json:"grievance_officer_details"`
	CustomerCareDetails     CustomerCareDetails     `json:"customer_care_details"`
	common_model.MessagingProduct
}
