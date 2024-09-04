package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	headers := http.Header{}
	headers.Set("Authorization", "ApiKey expected-key-value")

	apiKey, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expectedAPIKey := "expected-key-value"
	if apiKey != expectedAPIKey {
		t.Errorf("expected %v, got %v", expectedAPIKey, apiKey)
	}

}
