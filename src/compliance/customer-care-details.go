// Provides models to handle compliance services.
package compliance

type CustomerCareDetails struct {
	Email          string `json:"email"`
	LandlineNumber string `json:"landline_number"`
	MobileNumber   string `json:"mobile_number"`
}
