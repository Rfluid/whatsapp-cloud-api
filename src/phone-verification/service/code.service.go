package phone_verification_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

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
	formData := url.Values{}
	formData.Set("code_method", string(data.CodeMethod))
	formData.Set("language", data.Language)

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.RequestCode),
		bytes.NewBufferString(formData.Encode()),
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

// Verifies phone verification code.
func VerifyCode(
	api bootstrap_model.WhatsAppAPI,
	data phone_verification_model.VerifyCode,
) (common_model.SuccessResponse, error) {
	formData := url.Values{}
	formData.Set("code", string(data.Code))

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.VerifyCode),
		bytes.NewBufferString(formData.Encode()),
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
