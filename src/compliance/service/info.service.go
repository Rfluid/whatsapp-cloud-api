package compliance_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
	common_enum "github.com/Rfluid/whatsapp-cloud-api/src/common/enum"
	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
	compliance_model "github.com/Rfluid/whatsapp-cloud-api/src/compliance/model"
)

// Gets compliance info.
func Get(
	api bootstrap_model.WhatsAppAPI,
) (compliance_model.Info, error) {
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.BusinessComplianceInfo),
		nil,
	)
	req.Header = api.Headers

	resp, err := api.Client.Do(req)
	if err != nil {
		return compliance_model.Info{}, err
	}
	defer resp.Body.Close()

	var body compliance_model.Info

	json.NewDecoder(resp.Body).Decode(&body)

	return body, nil
}

// Creates a BusinessComplianceInfoSanitized (https://developers.facebook.com/docs/graph-api/reference/business-compliance-info-sanitized/).
func Post(
	api bootstrap_model.WhatsAppAPI,
	data compliance_model.PostInfoPayload,
) (common_model.SuccessResponse, error) {
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.BusinessComplianceInfo),
		bytes.NewBuffer(jsonData),
	)
	req.Header = api.Headers

	resp, err := api.Client.Do(req)
	if err != nil {
		return common_model.SuccessResponse{}, err
	}
	defer resp.Body.Close()

	var body common_model.SuccessResponse

	json.NewDecoder(resp.Body).Decode(&body)

	return body, nil
}
