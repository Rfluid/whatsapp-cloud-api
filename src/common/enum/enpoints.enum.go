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
)
