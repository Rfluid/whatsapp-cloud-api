package template_model

type Status string

const (
	Approved        Status = "APPROVED"
	InAppeal        Status = "IN_APPEAL"
	Pending         Status = "PENDING"
	Rejected        Status = "REJECTED"
	PendingDeletion Status = "PENDING_DELETION"
	Deleted         Status = "DELETED"
	Disabled        Status = "DISABLED"
	Paused          Status = "PAUSED"
	LimitExceeded   Status = "LIMIT_EXCEEDED"
	Archived        Status = "ARCHIVED"
)

// IsValid checks if the status is one of the allowed values.
func (s Status) IsValid() bool {
	switch s {
	case Approved, InAppeal, Pending, Rejected, PendingDeletion, Deleted, Disabled, Paused, LimitExceeded, Archived:
		return true
	default:
		return false
	}
}
