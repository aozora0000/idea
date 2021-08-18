package idea

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"time"
)

type Cipher struct {
	key []byte
}

func NewCipher(key []byte) Cipher {
	return Cipher{key: key}
}

func CreateNewKey(t time.Time) []byte {
	r := sha256.Sum256([]byte(t.String()))
	return r[:]
}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func (c Cipher) Encrypt(text []byte) (string, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", err
	}
	b := base64.StdEncoding.EncodeToString(text)
	cipt := make([]byte, aes.BlockSize+len(b))
	iv := cipt[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipt[aes.BlockSize:], []byte(b))
	return encodeBase64(cipt), nil
}

func (c Cipher) Decrypt(t string) (string, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", err
	}
	text, err := decodeBase64(t)
	if err != nil {
		return "", err
	}
	if len(text) < aes.BlockSize {
		return "", errors.New("too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	b, err := decodeBase64(string(text))
	if err != nil {
		return "", err
	}
	return string(b), nil
}
