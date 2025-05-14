// Provides two step verification models.
package two_step_verification_model

// Used to authenticate with two step verification.
type Pin struct {
	Pin string `json:"pin"`
}
