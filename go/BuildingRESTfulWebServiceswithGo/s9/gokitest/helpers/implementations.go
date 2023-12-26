package helpers

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

type EncryptServiceInstance struct{}

func (EncryptServiceInstance) Encrypt(_ context.Context, key string, plaintext string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	str := base64.StdEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(plaintext), nil))
	return str, nil
}

func (EncryptServiceInstance) Decrypt(_ context.Context, key string, cipherText string) (string, error) {
	cipherByte, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", fmt.Errorf("DecodeString: %w", err)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("aes.NewCipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("cipher.NewGCM: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(cipherByte) < nonceSize {
		return "", fmt.Errorf("invalid len of cipher text")
	}

	nonce, cipherByte := cipherByte[:nonceSize], cipherByte[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cipherByte, nil)
	if err != nil {
		return "", fmt.Errorf("gcm.open: %w", err)
	}

	return string(plaintext), nil
}
