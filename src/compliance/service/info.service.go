// Manipulates complicance info.
// Business compliance data is only exposed in the WhatsApp app and the WhatsApp Business app, and only if the app user's phone number is India-based (it begins with +91). So it's probably deprecated.
//
// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-account-to-number-current-status/business_compliance_info/
package compliance_service

import (
	"bytes"
	"encoding/json"
	"errors"
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
	req.Header = api.JSONHeaders

	resp, err := api.Client.Do(req)
	if err != nil {
		return compliance_model.Info{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errInt map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&errInt); err != nil {
			return compliance_model.Info{}, err
		}
		errMsgBytes, err := json.Marshal(errInt)
		if err != nil {
			return compliance_model.Info{}, err
		}
		return compliance_model.Info{}, errors.New(string(errMsgBytes))
	}

	var body compliance_model.Info

	json.NewDecoder(resp.Body).Decode(&body)

	return body, nil
}

// Creates a BusinessComplianceInfoSanitized (https://developers.facebook.com/docs/graph-api/reference/business-compliance-info-sanitized/).
func Post(
	api bootstrap_model.WhatsAppAPI,
	data compliance_model.PostInfoPayload,
) (common_model.SuccessResponse, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return common_model.SuccessResponse{}, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.BusinessComplianceInfo),
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
		var errInt map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&errInt); err != nil {
			return common_model.SuccessResponse{}, err
		}
		errMsgBytes, err := json.Marshal(errInt)
		if err != nil {
			return common_model.SuccessResponse{}, err
		}
		return common_model.SuccessResponse{}, errors.New(string(errMsgBytes))
	}

	var body common_model.SuccessResponse

	json.NewDecoder(resp.Body).Decode(&body)

	return body, nil
}
