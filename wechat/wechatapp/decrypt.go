package wechatapp

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
)

//摘自 https://github.com/medivhzhan/weapp

type watermark struct {
	AppID     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

type PhoneNumber struct {
	PhoneNumber     string    `json:"phoneNumber"`
	PurePhoneNumber string    `json:"purePhoneNumber"`
	CountryCode     string    `json:"countryCode"`
	Watermark       watermark `json:"watermark"`
}

func DecryptPhoneNumber(ssk, data, iv string) (phone PhoneNumber, err error) {
	bts, err := CBCDecrypt(ssk, data, iv)
	if err != nil {
		return
	}

	err = json.Unmarshal(bts, &phone)
	return
}

func CBCDecrypt(ssk, data, iv string) (bts []byte, err error) {
	key, err := base64.StdEncoding.DecodeString(ssk)
	if err != nil {
		return
	}

	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return
	}

	rawIV, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	size := aes.BlockSize

	if len(ciphertext) < size {
		err = errors.New("cipher too short")
		return
	}

	if len(ciphertext)%size != 0 {
		err = errors.New("cipher is not a multiple of the block size")
		return
	}

	mode := cipher.NewCBCDecrypter(block, rawIV[:size])
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	return PKCS5UnPadding(plaintext)
}

func PKCS5UnPadding(plaintext []byte) ([]byte, error) {
	ln := len(plaintext)

	unPadding := int(plaintext[ln-1])

	if unPadding > ln {
		return []byte{}, errors.New("数据不正确")
	}

	return plaintext[:(ln - unPadding)], nil
}
