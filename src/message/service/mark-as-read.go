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

// Marks message as read to user.
//
// One can only mark messages received as read. Use status as "read" to do so.
// After 30 days of message being received you now longer can mark it as read.
//
// See https://developers.facebook.com/docs/whatsapp/cloud-api/guides/mark-message-as-read/
func MarkAsRead(
	api bootstrap_model.WhatsAppAPI,
	data message_model.MarkAsRead,
) (common_model.SuccessResponse, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return common_model.SuccessResponse{}, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIDURL, common_enum.Messages),
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
