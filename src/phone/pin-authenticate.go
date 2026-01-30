// Performs two step verification with pin.
package phone

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Rfluid/whatsapp-cloud-api/src/bootstrap"
	"github.com/Rfluid/whatsapp-cloud-api/src/common"
)

// Authenticates with pin.
func AuthenticateWithPin(
	api bootstrap.WhatsAppAPI,
	pin Pin,
) (common.SuccessResponse, error) {
	jsonData, _ := json.Marshal(pin)

	req, err := http.NewRequest(
		"POST",
		api.WABAIDURL,
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return common.SuccessResponse{}, err
	}
	req.Header = api.JSONHeaders

	resp, err := api.Client.Do(req)
	if err != nil {
		return common.SuccessResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respErr common.ErrorResponse
		if decodeErr := json.NewDecoder(resp.Body).Decode(&respErr); decodeErr != nil {
			err = decodeErr
		} else {
			err = &respErr
		}
		return common.SuccessResponse{}, err
	}

	var body common.SuccessResponse

	err = json.NewDecoder(resp.Body).Decode(&body)

	return body, err
}
