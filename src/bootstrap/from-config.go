package bootstrap

import (
	"net/http"

)

// SenderConfig holds the parameters needed to build a fully configured WhatsAppAPI.
// Typically fetched from a database or configuration store.
type SenderConfig struct {
	AccessToken   string
	WABAID        string  // Phone number ID.
	WABAAccountID string  // WABA Account ID.
	Version       *string // Graph API version. Defaults to v24.0 if nil.
	CustomMainURL *string // Overrides the default graph API URL if set.
}

// FromConfig creates a ready-to-use WhatsAppAPI from a SenderConfig.
// A new http.Client is created for each call.
func FromConfig(cfg SenderConfig) (*WhatsAppAPI, error) {
	return FromConfigWithClient(cfg, &http.Client{})
}

// FromConfigWithClient creates a ready-to-use WhatsAppAPI from a SenderConfig
// using the provided http.Client. This allows sharing a single client across
// multiple senders for better connection pooling.
func FromConfigWithClient(cfg SenderConfig, client *http.Client) (*WhatsAppAPI, error) {
	api, err := GenerateWhatsAppAPI(cfg.AccessToken, cfg.Version, cfg.CustomMainURL)
	if err != nil {
		return nil, err
	}

	api.Client = client

	api, err = api.SetWABAID(cfg.WABAID)
	if err != nil {
		return nil, err
	}
	api.SetWABAIDURL(nil)

	api, err = api.SetWABAAccountID(cfg.WABAAccountID)
	if err != nil {
		return nil, err
	}
	api.SetWABAAccountIDURL(nil)

	api.SetJSONHeaders()
	api.SetFormHeaders()

	return api, nil
}
