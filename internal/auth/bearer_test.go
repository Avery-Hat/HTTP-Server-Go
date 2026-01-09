package auth

import (
	"net/http"
	"testing"
)

func TestGetBearerToken(t *testing.T) {
	h := make(http.Header)
	h.Set("Authorization", "Bearer abc123")

	got, err := GetBearerToken(h)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if got != "abc123" {
		t.Fatalf("expected abc123, got %q", got)
	}
}

func TestGetBearerToken_MissingHeader(t *testing.T) {
	h := make(http.Header)

	_, err := GetBearerToken(h)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestGetBearerToken_BadFormat(t *testing.T) {
	h := make(http.Header)
	h.Set("Authorization", "abc123")

	_, err := GetBearerToken(h)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
