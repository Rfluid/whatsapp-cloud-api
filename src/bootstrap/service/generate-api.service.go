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
) bootstrap_model.WhatsAppAPI {
	btp := bootstrap_model.WhatsAppAPI{
		Client: &http.Client{},
	}

	if version == nil {
		vrs := "v19.0"
		version = &vrs
	}

	mainUrl := "https://graph.facebook.com"
	mUrlP := &mainUrl

	btp.SetVersion(version).SetAccessToken(accessToken).SetMainURL(mUrlP, customMainURL)

	return btp
}
