// Provides two step verification models.
package phone

// Used to authenticate with two step verification.
type Pin struct {
	Pin string `json:"pin"`
}
