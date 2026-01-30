package typing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rfluid/whatsapp-cloud-api/src/bootstrap"
	"github.com/Rfluid/whatsapp-cloud-api/src/common"
)

// When you get a messages webhook indicating a received message,
// you can use the message.id value to mark the message as read
// and display a typing indicator so the WhatsApp user knows you
// are preparing a response. This is good practice if it will
// take you a few seconds to respond.
//
// The typing indicator will be dismissed once you respond,
// or after 25 seconds, whichever comes first.
// To prevent a poor user experience,
// only display a typing indicator if you are going to respond.
//
// See https://developers.facebook.com/docs/whatsapp/cloud-api/typing-indicators
func SendTyping(
	api bootstrap.WhatsAppAPI,
	data SendTypingPayload,
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
