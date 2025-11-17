package template_model

type TemplateCategory string

const (
	AccountUpdate         TemplateCategory = "ACCOUNT_UPDATE"
	PaymentUpdate         TemplateCategory = "PAYMENT_UPDATE"
	PersonalFinanceUpdate TemplateCategory = "PERSONAL_FINANCE_UPDATE"
	ShippingUpdate        TemplateCategory = "SHIPPING_UPDATE"
	ReservationUpdate     TemplateCategory = "RESERVATION_UPDATE"
	IssueResolution       TemplateCategory = "ISSUE_RESOLUTION"
	AppointmentUpdate     TemplateCategory = "APPOINTMENT_UPDATE"
	TransportationUpdate  TemplateCategory = "TRANSPORTATION_UPDATE"
	TicketUpdate          TemplateCategory = "TICKET_UPDATE"
	AlertUpdate           TemplateCategory = "ALERT_UPDATE"
	AutoReply             TemplateCategory = "AUTO_REPLY"
	Transactional         TemplateCategory = "TRANSACTIONAL"
	OTP                   TemplateCategory = "OTP"
	Utility               TemplateCategory = "UTILITY"
	Marketing             TemplateCategory = "MARKETING"
	Authentication        TemplateCategory = "AUTHENTICATION"
	FreeService           TemplateCategory = "FREE_SERVICE"
)

// IsValid checks if the category is one of the predefined values.
func (tc TemplateCategory) IsValid() bool {
	switch tc {
	case AccountUpdate,
		PaymentUpdate,
		PersonalFinanceUpdate,
		ShippingUpdate,
		ReservationUpdate,
		IssueResolution,
		AppointmentUpdate,
		TransportationUpdate,
		TicketUpdate,
		AlertUpdate,
		AutoReply,
		Transactional,
		OTP,
		Utility,
		Marketing,
		Authentication,
		FreeService:
		return true
	default:
		return false
	}
}
