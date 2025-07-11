package template_model

type Status string

const (
	Approved Status = "APPROVED"
	Pending  Status = "PENDING"
	Rejected Status = "REJECTED"
)

// IsValid checks if the status is one of the allowed values.
func (s Status) IsValid() bool {
	switch s {
	case Approved, Pending, Rejected:
		return true
	default:
		return false
	}
}
