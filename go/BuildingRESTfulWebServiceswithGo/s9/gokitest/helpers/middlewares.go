package helpers

import (
	"context"
	"time"

	"github.com/go-kit/log"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   EncryptService
}

func (mw LoggingMiddleware) Encrypt(ctx context.Context, key string, text string) (out string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "encryption",
			"key", key,
			"text", text,
			"output", out,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	out, err = mw.Next.Encrypt(ctx, key, text)
	return
}

func (mw LoggingMiddleware) Decrypt(ctx context.Context, key string, cipher string) (out string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "decryption",
			"key", key,
			"text", cipher,
			"output", out,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	out, err = mw.Next.Decrypt(ctx, key, cipher)
	return
}
