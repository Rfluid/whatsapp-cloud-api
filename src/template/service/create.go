package template_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
	common_enum "github.com/Rfluid/whatsapp-cloud-api/src/common/enum"
	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
	template_model "github.com/Rfluid/whatsapp-cloud-api/src/template/model"
)

// Creates template.
//
// Docs: https://developers.facebook.com/docs/whatsapp/business-management-api/message-templates/
func Create(
	api bootstrap_model.WhatsAppAPI,
	data template_model.CreateTemplate,
) (template_model.CreateTemplateResponse, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return template_model.CreateTemplateResponse{}, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAAccountID, common_enum.MessageTemplates),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return template_model.CreateTemplateResponse{}, err
	}
	req.Header = api.JSONHeaders

	resp, err := api.Client.Do(req)
	if err != nil {
		return template_model.CreateTemplateResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respErr common_model.ErrorResponse
		if decodeErr := json.NewDecoder(resp.Body).Decode(&respErr); decodeErr != nil {
			err = decodeErr
		} else {
			err = &respErr
		}
		return template_model.CreateTemplateResponse{}, err
	}

	var body template_model.CreateTemplateResponse

	err = json.NewDecoder(resp.Body).Decode(&body)

	return body, err
}
