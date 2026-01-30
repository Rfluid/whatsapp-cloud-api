// Queries the WhatsApp busines profile that your API is using.
package profile

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rfluid/whatsapp-cloud-api/src/bootstrap"
	"github.com/Rfluid/whatsapp-cloud-api/src/common"
)

// Gets WhatsApp Business Profile.
//
// @fields Are the fields that must be returned by query.
func GetProfile(
	api bootstrap.WhatsAppAPI,
	fields []BusinessProfileField,
) (BusinessProfile, error) {
	fieldsStr := ""

	if len(fields) > 0 {
		fieldsStr += string(fields[0])
		for _, field := range fields[1:] {
			fieldsStr += fmt.Sprintf(",%s", field)
		}
	}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/%s", api.WABAIDURL, common.BusinessProfile),
		nil,
	)
	if err != nil {
		return BusinessProfile{}, err
	}
	req.Header = api.JSONHeaders

	query := req.URL.Query()
	query.Add("fields", fieldsStr)
	req.URL.RawQuery = query.Encode()

	resp, err := api.Client.Do(req)
	if err != nil {
		return BusinessProfile{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respErr common.ErrorResponse
		if decodeErr := json.NewDecoder(resp.Body).Decode(&respErr); decodeErr != nil {
			err = decodeErr
		} else {
			err = &respErr
		}
		return BusinessProfile{}, err
	}

	var body BusinessProfile

	err = json.NewDecoder(resp.Body).Decode(&body)

	return body, err
}
