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
	customWABAIdURL *string,
) (bootstrap_model.WhatsAppAPI, error) {
	btp := bootstrap_model.WhatsAppAPI{
		Client: &http.Client{},
	}

	if version == nil {
		vrs := "v19.0"
		version = &vrs
	}

	mainUrl := "https://graph.facebook.com"
	mUrlP := &mainUrl

	_, err := btp.SetVersion(version).SetAccessToken(accessToken)
	if err != nil {
		return btp, err
	}
	btp.SetMainURL(mUrlP, customMainURL)

	return btp, nil
}
