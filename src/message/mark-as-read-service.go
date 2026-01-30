package message

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rfluid/whatsapp-cloud-api/src/bootstrap"
	"github.com/Rfluid/whatsapp-cloud-api/src/common"
)

// Marks message as read to user.
//
// One can only mark messages received as read. Use status as "read" to do so.
// After 30 days of message being received you now longer can mark it as read.
//
// See https://developers.facebook.com/docs/whatsapp/cloud-api/guides/mark-message-as-read/
func MarkAsRead(
	api bootstrap.WhatsAppAPI,
	data MarkAsReadPayload,
) (common.SuccessResponse, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return common.SuccessResponse{}, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIDURL, common.Messages),
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
