package ai

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/url"
	"strings"
)

var (
	signer = md5.New()
)

func GetRequestSign(value url.Values, key string) string {
	signer.Reset()
	io.WriteString(signer, value.Encode())
	io.WriteString(signer, "&app_key=")
	io.WriteString(signer, key)

	return strings.ToUpper(hex.EncodeToString(signer.Sum(nil)))
}
