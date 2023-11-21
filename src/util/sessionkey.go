package util

import "crypto/rand"
import "math/big"
import "encoding/base64"

func SessionKey() string {
	keyInt, err := rand.Int(rand.Reader, big.NewInt(4294967296))
	if err != nil {
		panic(err)
	}
	key := base64.RawURLEncoding.EncodeToString(keyInt.Bytes())
	return key
}
