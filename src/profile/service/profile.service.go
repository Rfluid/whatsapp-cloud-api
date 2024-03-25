package profile_service

import (
	"encoding/json"
	"fmt"
	"net/http"

	bootstrap_model "github.com/Rfluid/whatsapp/src/bootstrap/model"
	profile_model "github.com/Rfluid/whatsapp/src/profile/model"
)

// Gets WhatsApp Business Profile.
//
// @fields Are the fields that must be returned by query.
func GetProfile(
	api bootstrap_model.CloudApi,
	fields []profile_model.BusinessProfileFields,
) (profile_model.BusinessProfile, error) {
	fieldsStr := ""

	if len(fields) > 0 {
		fieldsStr += string(fields[0])
		for _, field := range fields[1:] {
			fieldsStr += fmt.Sprintf(", %s", field)
		}
	}

	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/whatsapp_business_profile", api.WABAIdURL),
		nil,
	)
	req.Header = api.Headers

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
