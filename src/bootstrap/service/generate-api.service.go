package bootstrap_service

import (
	"net/http"

	bootstrap_model "github.com/Rfluid/whatsapp/src/bootstrap/model"
)

func GenerateCloudAPI() bootstrap_model.CloudApi {
	return bootstrap_model.CloudApi{
		Prefix:  "https://graph.facebook.com",
		Version: "v18.0",
		Client:  &http.Client{},
	}
}
