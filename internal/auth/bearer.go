package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeader = errors.New("missing authorization header")
var ErrInvalidAuthHeader = errors.New("invalid authorization header")

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeader
	}

	const prefix = "ApiKey "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", ErrInvalidAuthHeader
	}

	key := strings.TrimSpace(strings.TrimPrefix(authHeader, prefix))
	if key == "" {
		return "", ErrInvalidAuthHeader
	}

	return key, nil
}

// changed bearer for new variables, commented out old func
func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeader
	}

	// Expected: "Bearer TOKEN"
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 {
		return "", ErrInvalidAuthHeader
	}

	if !strings.EqualFold(parts[0], "Bearer") {
		return "", ErrInvalidAuthHeader
	}

	token := strings.TrimSpace(parts[1])
	if token == "" {
		return "", ErrInvalidAuthHeader
	}

	return token, nil
}

// func GetBearerToken(headers http.Header) (string, error) {
// 	authHeader := headers.Get("Authorization")
// 	if authHeader == "" {
// 		return "", errors.New("missing authorization header")
// 	}

// 	// Expected: "Bearer TOKEN"
// 	parts := strings.SplitN(authHeader, " ", 2)
// 	if len(parts) != 2 {
// 		return "", errors.New("invalid authorization header format")
// 	}

// 	if !strings.EqualFold(parts[0], "Bearer") {
// 		return "", errors.New("authorization header is not bearer")
// 	}

// 	token := strings.TrimSpace(parts[1])
// 	if token == "" {
// 		return "", errors.New("missing bearer token")
// 	}

// 	return token, nil
// }
