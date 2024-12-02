package models

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"

	b64 "encoding/base64"
)

type Crypto struct {
	key []byte
}

func newCrypto(passkey string) Crypto {
	key := keyfy(passkey)

	return Crypto{key}
}

func (c Crypto) encryptB64(content string) (string, error) {
	res, err := c.encrypt(content)
	if err != nil {
		return "", err
	}

	return c.B64Encode(res), nil
}

func (c Crypto) decryptB64(content string) (string, error) {
	res, err := c.B64Decode(content)
	if err != nil {
		return "", err
	}

	return c.decrypt(res)
}

func (c Crypto) encrypt(content string) (string, error) {
	res, err := encrypt(content, c.key)

	return string(res), err
}

func (c Crypto) decrypt(content string) (string, error) {

	return decrypt([]byte(content), c.key)
}

func (c Crypto) B64Decode(content string) (string, error) {
	res, err := b64.StdEncoding.DecodeString(content)

	return string(res), err
}

func (c Crypto) B64Encode(content string) string {

	return b64.StdEncoding.EncodeToString([]byte(content))
}

func keyfy(password string) []byte {
	// this is to pad the password to make it 32
	password = fmt.Sprintf("%032s", password)
	// This is to avoid the problem that would arise
	// if characters of more than one byte were used.
	runes := []rune(password)

	key := ""
	for i := 0; i < 32; i++ {
		key += string(runes[i])
	}

	return []byte(key)
}

func encrypt(plaintext string, key []byte) ([]byte, error) {
	text := []byte(plaintext)
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, text, nil), nil
}

func decrypt(ciphertext []byte, key []byte) (string, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	bytes, err := gcm.Open(nil, nonce, ciphertext, nil)
	return string(bytes), err
}
