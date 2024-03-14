package utils

import (
	"log"
	"os"
	"time"
)

var (
	RedisChatExpiry = time.Since(time.Now().AddDate(0, 0, 1)) // 1 day
)

func GetEnv(key string) interface{} {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("env key %v not found", key)
	}
	return value
}

func GetSecret(key string) interface{} {
	secret := os.Getenv(key)
	// secret := cryptos.RSA_OAEP_Decrypt(GetEnv(key).(string), *RSA_PRIVATE_KEY)
	return secret
}
