package template_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
	common_enum "github.com/Rfluid/whatsapp-cloud-api/src/common/enum"
	common_service "github.com/Rfluid/whatsapp-cloud-api/src/common/service"
	template_model "github.com/Rfluid/whatsapp-cloud-api/src/template/model"
)

// Gets templates.
//
// @fields Are the fields that must be returned by query.
// @after Is used to paginate.
// @before Is used to paginate.
func Get(
	api bootstrap_model.WhatsAppAPI,
	fields []template_model.TemplateFields,
	after *string, // Used to paginate.
	before *string, // Used to paginate.
) (template_model.GetTemplateResponse, error) {
	var queryStringWg sync.WaitGroup

	queryCh := make(chan url.Values)

	fieldsStr := ""

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.MessageTemplates),
		nil,
	)
	if err != nil {
		return template_model.GetTemplateResponse{}, err
	}
	req.Header = api.JSONHeaders

	go func() {
		common_service.PropagateQueryParamsMultithread(
			queryCh,
			4,
			req,
		)
	}()

	queryStringWg.Add(1)
	go func() {
		defer queryStringWg.Done()

		if len(fields) > 0 {
			fieldsStr += string(fields[0])
			for _, field := range fields[1:] {
				fieldsStr += fmt.Sprintf(", %s", field)
				(<-queryCh).Add("fields", fieldsStr)
			}
		}
	}()

	queryStringWg.Add(1)
	go func() {
		defer queryStringWg.Done()
		if after != nil {
			(<-queryCh).Add("after", *after)
		}
	}()

	queryStringWg.Add(1)
	go func() {
		defer queryStringWg.Done()
		if before != nil {
			(<-queryCh).Add("before", *before)
		}
	}()

	queryStringWg.Wait()

	req.URL.RawQuery = (<-queryCh).Encode()

	close(queryCh)

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
