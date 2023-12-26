package helpers

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type EncryptService interface {
	Encrypt(ctx context.Context, key string, text string) (string, error)
	Decrypt(ctx context.Context, key string, cipher string) (string, error)
}

func MakeEncryptEndpoint(svc EncryptService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(EncryptRequest)
		msg, err := svc.Encrypt(ctx, req.Key, req.Text)
		if err != nil {
			return EncryptResponse{msg, err.Error()}, nil
		}
		return EncryptResponse{msg, ""}, nil
	}
}

func MakeDecryptEndpoint(svc EncryptService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DecryptRequest)
		msg, err := svc.Decrypt(ctx, req.Key, req.Cipher)
		if err != nil {
			return DecryptResponse{msg, err.Error()}, nil
		}
		return DecryptResponse{msg, ""}, nil
	}
}
