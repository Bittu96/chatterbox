package configs

import (
	"fmt"
	"log"
	utils "projects/chatterbox/server/pkgs/utilities"
	"reflect"
	"time"
)

type serverConfig struct {
	PORT                 string `env:"PORT" default:""`
	DB_Host              string `env:"DB_HOST" default:""`
	DB_Port              int    `env:"DB_PORT" default:""`
	DB_User              string `env:"DB_USER" default:""`
	DB_Pass              string `env:"DB_PASS" default:""`
	DB_Name              string `env:"DB_NAME" default:""`
	GOOGLE_CLIENT_ID     string `env:"GOOGLE_CLIENT_ID" default:""`
	GOOGLE_CLIENT_SECRET string `env:"GOOGLE_CLIENT_SECRET" default:""`
	REDIS_CHAT_EXPIRY    int    `env:"REDIS_CHAT_EXPIRY" default:""`
}

var ServerConfig serverConfig

var (
	RedisChatExpiry = time.Since(time.Now().AddDate(0, 0, ServerConfig.REDIS_CHAT_EXPIRY)) // 1 day
)

const (
	ENV     = "env"
	DEFAULT = "default"
)

func setConfig(destination serverConfig) {
	v := reflect.ValueOf(destination)
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get(ENV)

		if tag == "" || tag == "-" {
			continue
		}

		a := reflect.Indirect(reflect.ValueOf(destination))
		envVal := utils.GetEnv(tag)
		if envVal == "" {
			log.Fatal("Missing environment configuration for '" + a.Type().Field(i).Name + "', Loading default setting!")
		}

		switch t := v.Type().Field(i).Type.Kind(); t {
		case reflect.String:
			reflect.ValueOf(&destination).Elem().Field(i).SetString(envVal)
		case reflect.Int:
			reflect.ValueOf(&destination).Elem().Field(i).SetInt(int64(utils.ToInt(envVal)))
		}
	}

	ServerConfig = destination
}

func Load() {
	setConfig(ServerConfig)
	fmt.Println("server config:", ServerConfig)
}
