package cryptos

import (
	"encoding/base64"
	"log"
)

func B64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func B64Decode(str string) string {
	bStr, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}
	return string(bStr)
}
