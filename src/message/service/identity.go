package message_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
	common_enum "github.com/Rfluid/whatsapp-cloud-api/src/common/enum"
	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
	message_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model"
)

// Sets or unsets identity check.
//
// See https://developers.facebook.com/docs/whatsapp/cloud-api/reference/phone-numbers#verify
// to check examples of message with identity verification.
func ConfigIdentityCheck(
	api bootstrap_model.WhatsAppAPI,
	data message_model.UserIdentityChangeConfig,
) (common_model.SuccessResponse, error) {
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIDURL, common_enum.IdentitySettings),
		bytes.NewBuffer(jsonData),
	)
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
