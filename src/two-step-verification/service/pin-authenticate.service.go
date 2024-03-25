package two_step_verification_service

import (
	"bytes"
	"encoding/json"
	"net/http"

	bootstrap_model "github.com/Rfluid/whatsapp/src/bootstrap/model"
	common_model "github.com/Rfluid/whatsapp/src/common/model"
	two_step_verification_model "github.com/Rfluid/whatsapp/src/two-step-verification/model"
)

// Sends message.
func AuthenticateWithPin(
	api bootstrap_model.CloudApi,
	pin two_step_verification_model.Pin,
) (common_model.SuccessResponse, error) {
	jsonData, _ := json.Marshal(pin)

	req, _ := http.NewRequest(
		"POST",
		api.PhoneIdURL,
		bytes.NewBuffer(jsonData),
	)
	req.Header = api.Headers

	resp, err := api.Client.Do(req)
	if err != nil {
		return common_model.SuccessResponse{}, err
	}
	defer resp.Body.Close()

	var body common_model.SuccessResponse

	json.NewDecoder(resp.Body).Decode(&body)

	return body, nil
}
