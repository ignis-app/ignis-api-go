package util

import (
	"crypto/rand"
	"math/big"
	"encoding/base64"
)

func SessionKey() string {
	keyInt, keyErr := rand.Int(rand.Reader, big.NewInt(4294967296))
	if keyErr != nil {
		panic(keyErr)
	}
	key := base64.RawURLEncoding.EncodeToString(keyInt.Bytes())
	return key
}
