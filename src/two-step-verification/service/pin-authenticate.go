// Performs two step verification with pin.
package two_step_verification_service

import (
	"bytes"
	"encoding/json"
	"net/http"

	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
	two_step_verification_model "github.com/Rfluid/whatsapp-cloud-api/src/two-step-verification/model"
)

// Authenticates with pin.
func AuthenticateWithPin(
	api bootstrap_model.WhatsAppAPI,
	pin two_step_verification_model.Pin,
) (common_model.SuccessResponse, error) {
	jsonData, _ := json.Marshal(pin)

	req, err := http.NewRequest(
		"POST",
		api.WABAIdURL,
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return common_model.SuccessResponse{}, err
	}
	req.Header = api.JSONHeaders

	resp, err := api.Client.Do(req)
	if err != nil {
		return common_model.SuccessResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respErr common_model.ErrorResponse
		if decodeErr := json.NewDecoder(resp.Body).Decode(&respErr); decodeErr != nil {
			err = decodeErr
		} else {
			err = &respErr
		}
		return common_model.SuccessResponse{}, err
	}

	var body common_model.SuccessResponse

	err = json.NewDecoder(resp.Body).Decode(&body)

	return body, err
}
