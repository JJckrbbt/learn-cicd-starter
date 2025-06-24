package auth

import (
	"net/http/httptest"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	reqValid := httptest.NewRequest("GET", "/some/api/route", nil)
	reqValid.Header.Set("Authorization", "ApiKey valid-api-key")

	apiKey, err := GetAPIKey(reqValid.Header)
	if err != nil {
		t.Errorf("handler err: %v", err)
	}

	if apiKey != "valid-api-key" {
		t.Errorf("handler returned wrong key: %s", apiKey)
	}

	reqInvalid := httptest.NewRequest("GET", "/some/api/route", nil)
	reqInvalid.Header.Set("Authorization", "apikey invalid-api-key")

	apiKey, err = GetAPIKey(reqInvalid.Header)
	if err == nil {
		t.Errorf("handler did not err: %v", err)
	}

	if apiKey == "invalid-api-key" {
		t.Errorf("handler returned validated wrong key: %s", apiKey)
	}

	reqNoAuthHeader := httptest.NewRequest("GET", "/some/api/route", nil)
	reqNoAuthHeader.Header.Set("NotAnAuthorization", "ApiKey no-auth-header")

	_, err = GetAPIKey(reqNoAuthHeader.Header)
	if err == nil {
		t.Errorf("handler returned validated key in error")
	}

}
