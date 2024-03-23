package bootstrap_model

import (
	"fmt"
	"log"
)

type ApiBootstrap struct {
	AccessToken string
	Prefix      string
	Version     string
	MainURL     string
}

func (btp *ApiBootstrap) SetMainURL() {
	btp.MainURL = fmt.Sprintf("%s/%s", btp.Prefix, btp.Version)
}

func (btp *ApiBootstrap) SetAccessToken(
	accessToken string,
) *ApiBootstrap {
	if accessToken != "" {
		btp.AccessToken = accessToken
	} else {
		log.Fatalln("Access token was not provided.")
	}
	return btp
}

func (btp *ApiBootstrap) SetVersion(
	version string,
) *ApiBootstrap {
	if version != "" {
		btp.Version = version
	} else {
		log.Println("Version not provided. Loading default...")
	}
	return btp
}

func (btp *ApiBootstrap) SetPrefix(
	prefix string,
) *ApiBootstrap {
	if prefix != "" {
		btp.Prefix = prefix
	} else {
		log.Println("Prefix not provided. Loading default...")
	}

	return btp
}
