package helpers

import (
	"context"
	"encoding/json"
	"net/http"
)

func DecodeEncriptRequest(_ context.Context, r *http.Request) (any, error) {
	var req EncryptRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeDecryptRequest(_ context.Context, r *http.Request) (any, error) {
	var req DecryptRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response any) error {
	return json.NewEncoder(w).Encode(response)
}
