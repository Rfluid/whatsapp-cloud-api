package template

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rfluid/whatsapp-cloud-api/src/bootstrap"
	"github.com/Rfluid/whatsapp-cloud-api/src/common"
)

// Creates 
//
// Docs: https://developers.facebook.com/docs/whatsapp/business-management-api/message-templates/
func Create(
	api bootstrap.WhatsAppAPI,
	data CreateTemplate,
) (CreateTemplateResponse, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return CreateTemplateResponse{}, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAAccountID, common.MessageTemplates),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return CreateTemplateResponse{}, err
	}
	req.Header = api.JSONHeaders

	resp, err := api.Client.Do(req)
	if err != nil {
		return CreateTemplateResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respErr common.ErrorResponse
		if decodeErr := json.NewDecoder(resp.Body).Decode(&respErr); decodeErr != nil {
			err = decodeErr
		} else {
			err = &respErr
		}
		return CreateTemplateResponse{}, err
	}

	var body CreateTemplateResponse

	err = json.NewDecoder(resp.Body).Decode(&body)

	return body, err
}
