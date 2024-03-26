package bootstrap_model

import (
	"fmt"
	"log"
	"net/http"
)

type WhatsAppAPI struct {
	AccessToken string
	WABAId      string
	Version     string
	MainURL     string
	WABAIdURL   string
	Headers     http.Header
	Client      *http.Client
}

func (btp *WhatsAppAPI) SetWABAIdURL(customUrl *string) {
	if customUrl != nil {
		btp.WABAIdURL = *customUrl
	} else {
		btp.WABAIdURL = fmt.Sprintf("%s/%s", btp.MainURL, btp.WABAId)
	}
}

func (btp *WhatsAppAPI) SetAccessToken(
	accessToken string,
) *WhatsAppAPI {
	if accessToken != "" {
		btp.AccessToken = accessToken
	} else {
		log.Fatalln("Access token was not provided.")
	}
	return btp
}

func (btp *WhatsAppAPI) SetWABAId(
	waba string,
) *WhatsAppAPI {
	if waba != "" {
		btp.WABAId = waba
	} else {
		log.Fatalln("WABA id was not provided.")
	}
	return btp
}

func (btp *WhatsAppAPI) SetVersion(
	version *string,
) *WhatsAppAPI {
	if version != nil {
		btp.Version = *version
	}
	// else {
	// 	log.Println("Version not provided. Loading default...")
	// }
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
	// else {
	// 	log.Println("MainURL not provided. Loading default...")
	// }

	return btp
}

func (btp *WhatsAppAPI) SetHeaders() *WhatsAppAPI {
	btp.Headers = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {fmt.Sprintf("Bearer %s", btp.AccessToken)},
	}
	return btp
}
