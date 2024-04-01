package message_service

import (
	"bytes"
	"encoding/json"
	"errors"
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
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.IdentitySettings),
		bytes.NewBuffer(jsonData),
	)
	req.Header = api.JSONHeaders

	resp, err := api.Client.Do(req)
	if err != nil {
		return common_model.SuccessResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errInt map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&errInt); err != nil {
			return common_model.SuccessResponse{}, err
		}
		errMsgBytes, err := json.MarshalIndent(errInt, "", "    ")
		if err != nil {
			return common_model.SuccessResponse{}, err
		}
		return common_model.SuccessResponse{}, errors.New(string(errMsgBytes))
	}

	var body common_model.SuccessResponse

	json.NewDecoder(resp.Body).Decode(&body)

	return body, nil
}
