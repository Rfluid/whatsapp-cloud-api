// Useful functions to bootstrap the API.
package bootstrap_service

import (
	"net/http"

	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
)

func GenerateWhatsAppAPI(
	accessToken string,
	version *string,
	customMainURL *string,
) (*bootstrap_model.WhatsAppAPI, error) {
	btp := &bootstrap_model.WhatsAppAPI{
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
