package profile_service

import (
	"encoding/json"
	"fmt"
	"net/http"

	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
	common_enum "github.com/Rfluid/whatsapp-cloud-api/src/common/enum"
	profile_model "github.com/Rfluid/whatsapp-cloud-api/src/profile/model"
)

// Gets WhatsApp Business Profile.
//
// @fields Are the fields that must be returned by query.
func GetProfile(
	api bootstrap_model.WhatsAppAPI,
	fields []profile_model.BusinessProfileField,
) (profile_model.BusinessProfile, error) {
	fieldsStr := ""

	if len(fields) > 0 {
		fieldsStr += string(fields[0])
		for _, field := range fields[1:] {
			fieldsStr += fmt.Sprintf(", %s", field)
		}
	}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.BusinessProfile),
		nil,
	)
	if err != nil {
		return profile_model.BusinessProfile{}, err
	}
	req.Header = api.JSONHeaders

	query := req.URL.Query()
	query.Add("fields", fieldsStr)
	req.URL.RawQuery = query.Encode()

	resp, err := api.Client.Do(req)
	if err != nil {
		return profile_model.BusinessProfile{}, err
	}
	defer resp.Body.Close()

	var body profile_model.BusinessProfile

	json.NewDecoder(resp.Body).Decode(&body)

	return body, nil
}
