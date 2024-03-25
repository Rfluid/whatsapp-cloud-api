package message_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	bootstrap_model "github.com/Rfluid/whatsapp/src/bootstrap/model"
	common_enum "github.com/Rfluid/whatsapp/src/common/enum"
	message_model "github.com/Rfluid/whatsapp/src/message/model"
)

// Changes message status.
func SetStatus(
	api bootstrap_model.CloudApi,
	data message_model.MessageStatus,
) (message_model.Response, error) {
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.Messages),
		bytes.NewBuffer(jsonData),
	)
	req.Header = api.Headers

	resp, err := api.Client.Do(req)
	if err != nil {
		return message_model.Response{}, err
	}
	defer resp.Body.Close()

	var body message_model.Response

	json.NewDecoder(resp.Body).Decode(&body)

	return body, nil
}
