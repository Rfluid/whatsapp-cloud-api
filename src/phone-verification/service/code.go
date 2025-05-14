// Handles phone validation with code.
package phone_verification_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
	common_enum "github.com/Rfluid/whatsapp-cloud-api/src/common/enum"
	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
	phone_verification_model "github.com/Rfluid/whatsapp-cloud-api/src/phone-verification/model"
)

// Requests phone verification code.
func RequestCode(
	api bootstrap_model.WhatsAppAPI,
	data phone_verification_model.RequestCode,
) (common_model.SuccessResponse, error) {
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.RequestCode),
		bytes.NewBufferString(data.ToURLValues().Encode()),
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

// Verifies phone verification code.
func VerifyCode(
	api bootstrap_model.WhatsAppAPI,
	data phone_verification_model.VerifyCode,
) (common_model.SuccessResponse, error) {
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.VerifyCode),
		bytes.NewBufferString(data.ToURLValues().Encode()),
	)
	if err != nil {
		return common_model.SuccessResponse{}, err
	}

	req.Header = api.FormHeaders

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
