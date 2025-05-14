package webhook_model

// Field affected at webhook call.
//
// https://developers.facebook.com/docs/whatsapp/business-management-api/guides/set-up-webhooks
// https://developers.facebook.com/docs/graph-api/webhooks/reference/whatsapp-business-account/
type Field string

const (
	AccountAlerts                Field = "account_alerts"
	BusinessCapabilityUpdate     Field = "business_capability_update"
	BusinessStatusUpdate         Field = "business_status_update"
	CampaignStatusUpdate         Field = "campaign_status_update"
	AccountUpdate                Field = "account_update"
	Security                     Field = "security"
	AccountReviewUpdate          Field = "account_review_update"
	PhoneNumberQualityUpdate     Field = "phone_number_quality_update"
	PhoneNumberNameUpdate        Field = "phone_number_name_update"
	PartnerSolutions             Field = "partner_solutions"
	MessagingHandovers           Field = "messaging_handovers"
	Flows                        Field = "flows"
	MessageEchoes                Field = "message_echoes"
	TemplateCategoryUpdate       Field = "template_category_update"
	MessageTemplateQualityUpdate Field = "message_template_quality_update"
	MessageTemplateStatusUpdate  Field = "message_template_status_update"
	Messages                     Field = "messages"
)
