// Handles number registering to WhatsApp.
package phone

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rfluid/whatsapp-cloud-api/src/bootstrap"
	"github.com/Rfluid/whatsapp-cloud-api/src/common"
)

// Register WhatsApp business number.
func Register(
	api bootstrap.WhatsAppAPI,
	data RegisterPayload,
) (common.SuccessResponse, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return common.SuccessResponse{}, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIDURL, common.Register),
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

// Deregister WhatsApp business number.
func DeRegister(
	api bootstrap.WhatsAppAPI,
) (common.SuccessResponse, error) {
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIDURL, common.DeRegister),
		nil,
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
