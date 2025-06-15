package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extracts api key from headers of request

// ex:
// Authorization : ApiKey {insert api key here}
func GetAPIKey(headers http.Header) (string, error) {
	// apikey blhablahbigapikey
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication header found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed Auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return vals[1], nil
}
