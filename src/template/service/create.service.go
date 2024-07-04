package template_service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
	common_enum "github.com/Rfluid/whatsapp-cloud-api/src/common/enum"
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
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.MessageTemplates),
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
		var errInt map[string]interface{}

		if err := json.NewDecoder(resp.Body).Decode(&errInt); err != nil {
			return template_model.CreateTemplateResponse{}, err
		}

		errMsgBytes, err := json.Marshal(errInt)
		if err != nil {
			return template_model.CreateTemplateResponse{}, err
		}

		return template_model.CreateTemplateResponse{}, errors.New(string(errMsgBytes))
	}

	var body template_model.CreateTemplateResponse

	json.NewDecoder(resp.Body).Decode(&body)

	return body, nil
}
