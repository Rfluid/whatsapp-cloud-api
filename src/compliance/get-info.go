// Manipulates complicance info.
// Business compliance data is only exposed in the WhatsApp app and the WhatsApp Business app, and only if the app user's phone number is India-based (it begins with +91). So it's probably deprecated.
//
// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-account-to-number-current-status/business_compliance_info/
package compliance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rfluid/whatsapp-cloud-api/src/bootstrap"
	"github.com/Rfluid/whatsapp-cloud-api/src/common"
)

// Gets compliance info.
func Get(
	api bootstrap.WhatsAppAPI,
) (Info, error) {
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/%s", api.WABAIDURL, common.BusinessComplianceInfo),
		nil,
	)
	req.Header = api.JSONHeaders

	resp, err := api.Client.Do(req)
	if err != nil {
		return Info{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respErr common.ErrorResponse
		if decodeErr := json.NewDecoder(resp.Body).Decode(&respErr); decodeErr != nil {
			err = decodeErr
		} else {
			err = &respErr
		}
		return Info{}, err
	}

	var body Info

	err = json.NewDecoder(resp.Body).Decode(&body)

	return body, err
}

// Creates a BusinessComplianceInfoSanitized (https://developers.facebook.com/docs/graph-api/reference/business-compliance-info-sanitized/).
func Post(
	api bootstrap.WhatsAppAPI,
	data PostInfoPayload,
) (common.SuccessResponse, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return common.SuccessResponse{}, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIDURL, common.BusinessComplianceInfo),
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
