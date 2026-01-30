package bootstrap

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFromConfig_Success(t *testing.T) {
	cfg := SenderConfig{
		AccessToken:   "test-token",
		WABAID:        "123456",
		WABAAccountID: "abcdef",
	}

	api, err := FromConfig(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if api.AccessToken != "test-token" {
		t.Errorf("AccessToken = %q, want %q", api.AccessToken, "test-token")
	}
	if api.WABAID != "123456" {
		t.Errorf("WABAID = %q, want %q", api.WABAID, "123456")
	}
	if api.WABAAccountID != "abcdef" {
		t.Errorf("WABAAccountID = %q, want %q", api.WABAAccountID, "abcdef")
	}
	if api.Version != "v24.0" {
		t.Errorf("Version = %q, want %q", api.Version, "v24.0")
	}

	expectedWABAIDURL := fmt.Sprintf("%s/%s", api.MainURL, "123456")
	if api.WABAIDURL != expectedWABAIDURL {
		t.Errorf("WABAIDURL = %q, want %q", api.WABAIDURL, expectedWABAIDURL)
	}

	expectedAccountURL := fmt.Sprintf("%s/%s", api.MainURL, "abcdef")
	if api.WABAAccountIDURL != expectedAccountURL {
		t.Errorf("WABAAccountIDURL = %q, want %q", api.WABAAccountIDURL, expectedAccountURL)
	}

	if api.JSONHeaders.Get("Content-Type") != "application/json" {
		t.Error("JSONHeaders missing Content-Type")
	}
	if api.JSONHeaders.Get("Authorization") != "Bearer test-token" {
		t.Error("JSONHeaders missing Authorization")
	}
	if api.FormHeaders.Get("Authorization") != "Bearer test-token" {
		t.Error("FormHeaders missing Authorization")
	}
	if api.Client == nil {
		t.Error("Client should not be nil")
	}
}

func TestFromConfig_CustomVersion(t *testing.T) {
	v := "v19.0"
	cfg := SenderConfig{
		AccessToken:   "tok",
		WABAID:        "1",
		WABAAccountID: "2",
		Version:       &v,
	}

	api, err := FromConfig(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if api.Version != "v19.0" {
		t.Errorf("Version = %q, want %q", api.Version, "v19.0")
	}
}

func TestFromConfig_EmptyAccessToken(t *testing.T) {
	cfg := SenderConfig{
		WABAID:        "1",
		WABAAccountID: "2",
	}

	_, err := FromConfig(cfg)
	if err == nil {
		t.Fatal("expected error for empty access token")
	}
}

func TestFromConfig_EmptyWABAID(t *testing.T) {
	cfg := SenderConfig{
		AccessToken:   "tok",
		WABAAccountID: "2",
	}

	_, err := FromConfig(cfg)
	if err == nil {
		t.Fatal("expected error for empty WABA ID")
	}
}

func TestFromConfig_EmptyWABAAccountID(t *testing.T) {
	cfg := SenderConfig{
		AccessToken: "tok",
		WABAID:      "1",
	}

	_, err := FromConfig(cfg)
	if err == nil {
		t.Fatal("expected error for empty WABA Account ID")
	}
}

func TestFromConfigWithClient_SharedClient(t *testing.T) {
	client := &http.Client{}

	cfg1 := SenderConfig{AccessToken: "t1", WABAID: "1", WABAAccountID: "a1"}
	cfg2 := SenderConfig{AccessToken: "t2", WABAID: "2", WABAAccountID: "a2"}

	api1, err := FromConfigWithClient(cfg1, client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	api2, err := FromConfigWithClient(cfg2, client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if api1.Client != client || api2.Client != client {
		t.Error("both APIs should share the same http.Client")
	}
}
