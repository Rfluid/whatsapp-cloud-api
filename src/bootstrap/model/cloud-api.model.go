// Provides modules to bootstrap API.
package bootstrap_model

import (
	"fmt"
	"net/http"
)

type WhatsAppAPI struct {
	AccessToken string
	WABAId      string
	Version     string
	MainURL     string
	WABAIdURL   string
	JSONHeaders http.Header
	FormHeaders http.Header
	Client      *http.Client
}

func (btp *WhatsAppAPI) SetWABAIdURL(customUrl *string) {
	if customUrl != nil {
		btp.WABAIdURL = *customUrl
		return
	}
	btp.WABAIdURL = fmt.Sprintf("%s/%s", btp.MainURL, btp.WABAId)
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
