package models

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

type Crypto struct {
	key []byte
}

func newCrypto(passkey string) Crypto {
	key := keyfy(passkey)

	return Crypto{key}
}

/** test method **
func (c Crypto) GetKey() string {
	return strings.Trim(string(c.key), "0")
}
*/

func (c Crypto) encryptB64(content string) (string, error) {
	res, err := c.encrypt(content)
	if err != nil {
		return "", err
	}

	return c.b64Encode(res), nil
}

func (c Crypto) decryptB64(content string) (string, error) {
	res, err := c.b64Decode(content)
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

func (c Crypto) b64Decode(content string) (string, error) {
	res, err := base64.StdEncoding.DecodeString(content)

	return string(res), err
}

func (c Crypto) b64Encode(content string) string {

	return base64.StdEncoding.EncodeToString([]byte(content))
}

/*
// Verifies that a string contains only ASCII characters (single byte):
// https://stackoverflow.com/questions/53069040/checking-a-string-contains-only-ascii-characters
func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}

	return true
}
*/

func keyfy(password string) []byte {
	// this is to pad the password to make it 32
	password = fmt.Sprintf("%032s", password)

	key := ""
	for i := 0; i < 32; i++ {
		key += string(password[i])
	}

	return []byte(key)
}

func encrypt(plaintext string, key []byte) ([]byte, error) {
	aes, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return nil, err
	}

	// We need a 12-byte nonce for GCM (modifiable if you use cipher.NewGCMWithNonceSize())
	// A nonce should always be randomly generated for every encryption.
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	// https://cs.opensource.google/go/go/+/refs/tags/go1.23.4:src/crypto/rand/rand.go;l=26
	if err != nil {
		return nil, err
	}

	// ciphertext here is actually nonce+ciphertext
	// So that when we decrypt, just knowing the nonce size
	// is enough to separate it from the ciphertext.
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	return ciphertext, nil
}

func decrypt(ciphertext []byte, key []byte) (string, error) {
	aes, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return "", err
	}

	// Since we know the ciphertext is actually nonce+ciphertext
	// And len(nonce) == NonceSize(). We can separate the two.
	// Therefore, we are sure that len(ciphertext) < nonceSize
	// will never be true.
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

/* REFERENCES:
https://dev.to/breda/secret-key-encryption-with-go-using-aes-316d
https://www.geeksforgeeks.org/io-readfull-function-in-golang-with-examples/
https://blog.logrocket.com/learn-golang-encryption-decryption/
https://www.twilio.com/en-us/blog/encrypt-and-decrypt-data-in-go-with-aes-256
https://bruinsslot.jp/post/golang-crypto/

https://www.google.com/search?q=golang+password+encryption+and+decryption&oq=golang+password+encryption+and+decryption&gs_lcrp=EgZjaHJvbWUqBggAEEUYOzIGCAAQRRg7MgYIARBFGEAyCggCEAAYgAQYogQyCggDEAAYgAQYogQyCggEEAAYgAQYogQyCggFEAAYgAQYogTSAQgxODI5ajBqN6gCCLACAQ&sourceid=chrome&ie=UTF-8
*/
