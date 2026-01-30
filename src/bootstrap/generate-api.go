// Useful functions to bootstrap the API.
package bootstrap

import (
	"net/http"

)

func GenerateWhatsAppAPI(
	accessToken string,
	version *string,
	customMainURL *string,
) (*WhatsAppAPI, error) {
	btp := &WhatsAppAPI{
		Client: &http.Client{},
	}

	if version == nil {
		vrs := "v24.0"
		version = &vrs
	}

	mainUrl := "https://graph.facebook.com"
	mUrlP := &mainUrl

	var err error
	btp, err = btp.SetVersion(version).SetAccessToken(accessToken)
	if err != nil {
		return btp, err
	}

	btp.SetMainURL(mUrlP, customMainURL)

	return btp, nil
}
