package template_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
	common_enum "github.com/Rfluid/whatsapp-cloud-api/src/common/enum"
	template_model "github.com/Rfluid/whatsapp-cloud-api/src/template/model"
)

// Gets templates.
//
// @fields Are the fields that must be returned by query.
// @query Is mainly used to paginate.
func Get(
	api bootstrap_model.WhatsAppAPI,
	query template_model.GetQueryParams,
) (template_model.GetTemplateResponse, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.MessageTemplates),
		nil,
	)
	if err != nil {
		return template_model.GetTemplateResponse{}, err
	}

	query.BuildQuery(req)

	resp, err := api.Client.Do(req)
	if err != nil {
		return template_model.GetTemplateResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errInt map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&errInt); err != nil {
			return template_model.GetTemplateResponse{}, err
		}
		errMsgBytes, err := json.MarshalIndent(errInt, "", "    ")
		if err != nil {
			return template_model.GetTemplateResponse{}, err
		}
		return template_model.GetTemplateResponse{}, errors.New(string(errMsgBytes))
	}

	var body template_model.GetTemplateResponse

	json.NewDecoder(resp.Body).Decode(&body)

	return body, nil
}
