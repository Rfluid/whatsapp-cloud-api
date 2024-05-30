package template_model

type Status string

const (
	Approved Status = "APPROVED"
	Pending  Status = "PENDING"
	Rejected Status = "REJECTED"
)
