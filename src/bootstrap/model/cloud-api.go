// Provides modules to bootstrap API.
package bootstrap_model

import (
	"fmt"
	"net/http"
)

type WhatsAppAPI struct {
	AccessToken      string       // Access token that is combined with Bearer to authenticate requests.
	WABAId           string       // Number of ID.
	WABAAccountId    string       // WABA Account ID. Available at https://developers.facebook.com/apps/APP_ID/whatsapp-business/wa-dev-console/?business_id=BUSINESS_ID
	Version          string       // Verision of Meta's graph API.
	MainURL          string       // Meta's graph API URL with version.
	WABAIdURL        string       // Used to send messages and things related to number.
	WABAAccountIdURL string       // Used to get account info such as templates, number, etc.
	JSONHeaders      http.Header  // Authenticates JSON requests.
	FormHeaders      http.Header  // Authenticates form requests.
	Client           *http.Client // Client to make requests.
}

func (btp *WhatsAppAPI) SetWABAIdURL(customUrl *string) {
	if customUrl != nil {
		btp.WABAIdURL = *customUrl
		return
	}
	btp.WABAIdURL = fmt.Sprintf("%s/%s", btp.MainURL, btp.WABAId)
}

func (btp *WhatsAppAPI) SetWABAAccountIdURL(customUrl *string) {
	if customUrl != nil {
		btp.WABAAccountIdURL = *customUrl
		return
	}
	btp.WABAAccountIdURL = fmt.Sprintf("%s/%s", btp.MainURL, btp.WABAAccountId)
}

func (btp *WhatsAppAPI) SetAccessToken(
	accessToken string,
) (*WhatsAppAPI, error) {
	if accessToken == "" {
		return nil, fmt.Errorf("access token was not provided")
	}
	btp.AccessToken = accessToken
	return btp, nil
}

func (btp *WhatsAppAPI) SetWABAId(
	waba string,
) (*WhatsAppAPI, error) {
	if waba == "" {
		return nil, fmt.Errorf("WABA ID was not provided")
	}
	btp.WABAId = waba

	return btp, nil
}

func (btp *WhatsAppAPI) SetWABAAccountId(
	accountId string,
) (*WhatsAppAPI, error) {
	if accountId == "" {
		return nil, fmt.Errorf("WABA Account ID was not provided")
	}
	btp.WABAAccountId = accountId

	return btp, nil
}

func (btp *WhatsAppAPI) SetVersion(
	version *string,
) *WhatsAppAPI {
	if version != nil {
		btp.Version = *version
	}
	return btp
}

func (btp *WhatsAppAPI) SetMainURL(
	prefix *string,
	customUrl *string,
) *WhatsAppAPI {
	if customUrl != nil {
		btp.MainURL = *customUrl
	} else if prefix != nil {
		btp.MainURL = fmt.Sprintf("%s/%s", *prefix, btp.Version)
	}

	return btp
}

func (btp *WhatsAppAPI) SetJSONHeaders() *WhatsAppAPI {
	btp.JSONHeaders = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {fmt.Sprintf("Bearer %s", btp.AccessToken)},
	}
	return btp
}

func (btp *WhatsAppAPI) SetFormHeaders() *WhatsAppAPI {
	btp.FormHeaders = http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", btp.AccessToken)},
	}
	return btp
}
