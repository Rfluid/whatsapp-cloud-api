package phone_verification_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	bootstrap_model "github.com/Rfluid/whatsapp/src/bootstrap/model"
	common_enum "github.com/Rfluid/whatsapp/src/common/enum"
	common_model "github.com/Rfluid/whatsapp/src/common/model"
	phone_verification_model "github.com/Rfluid/whatsapp/src/phone-verification/model"
)

// Sets or unsets identity check.
//
// See https://developers.facebook.com/docs/whatsapp/cloud-api/reference/phone-numbers#verify
// to check examples of message with identity verification.
func ConfigIdentityCheck(
	api bootstrap_model.CloudApi,
	data phone_verification_model.UserIdentityChangeConfig,
) (common_model.SuccessResponse, error) {
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.IdentitySettings),
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
