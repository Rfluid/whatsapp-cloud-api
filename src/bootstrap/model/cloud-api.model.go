package bootstrap_model

import (
	"fmt"
	"log"
	"net/http"
)

type CloudApi struct {
	AccessToken string
	Prefix      string
	Version     string
	WABAId      string
	MainURL     string
	Headers     http.Header
	Client      *http.Client
}

func (btp *CloudApi) SetMainURL() {
	btp.MainURL = fmt.Sprintf("%s/%s/%s", btp.Prefix, btp.Version, btp.WABAId)
}

func (btp *CloudApi) SetAccessToken(
	accessToken string,
) *CloudApi {
	if accessToken != "" {
		btp.AccessToken = accessToken
	} else {
		log.Fatalln("Access token was not provided.")
	}
	return btp
}

func (btp *CloudApi) SetWABAId(
	waba string,
) *CloudApi {
	if waba != "" {
		btp.WABAId = waba
	} else {
		log.Fatalln("Access token was not provided.")
	}
	return btp
}

func (btp *CloudApi) SetVersion(
	version string,
) *CloudApi {
	if version != "" {
		btp.Version = version
	} else {
		log.Println("Version not provided. Loading default...")
	}
	return btp
}

func (btp *CloudApi) SetPrefix(
	prefix string,
) *CloudApi {
	if prefix != "" {
		btp.Prefix = prefix
	} else {
		log.Println("Prefix not provided. Loading default...")
	}

	return btp
}

func (btp *CloudApi) SetHeaders() *CloudApi {
	btp.Headers = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {fmt.Sprintf("Bearer %s", btp.AccessToken)},
	}
	return btp
}
