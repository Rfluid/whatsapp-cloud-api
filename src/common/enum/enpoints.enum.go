// Provides standard enums used throughout the application.
package common_enum

type Endpoint string

const (
	Messages               Endpoint = "messages"
	BusinessComplianceInfo Endpoint = "business_compliance_info"
	RequestCode            Endpoint = "request_code"
	VerifyCode             Endpoint = "verify_code"
	IdentitySettings       Endpoint = "settings"
	BusinessProfile        Endpoint = "whatsapp_business_profile"
	Register               Endpoint = "register"
	DeRegister             Endpoint = "deregister"
	Media                  Endpoint = "media"
	MessageTemplates       Endpoint = "message_templates"
)
