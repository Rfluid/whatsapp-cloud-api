// Handles phone validation with code.
package phone

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rfluid/whatsapp-cloud-api/src/bootstrap"
	"github.com/Rfluid/whatsapp-cloud-api/src/common"
)

// Requests phone verification code.
func RequestCode(
	api bootstrap.WhatsAppAPI,
	data RequestCodePayload,
) (common.SuccessResponse, error) {
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIDURL, common.RequestCode),
		bytes.NewBufferString(data.ToURLValues().Encode()),
	)
	if err != nil {
		return common.SuccessResponse{}, err
	}

	req.Header = api.FormHeaders

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

// Verifies phone verification code.
func VerifyCode(
	api bootstrap.WhatsAppAPI,
	data VerifyCodePayload,
) (common.SuccessResponse, error) {
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIDURL, common.VerifyCode),
		bytes.NewBufferString(data.ToURLValues().Encode()),
	)
	if err != nil {
		return common.SuccessResponse{}, err
	}

	req.Header = api.FormHeaders

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
