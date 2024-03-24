package secrets

import (
	"fmt"
	utils "projects/chatterbox/server/pkgs/utilities"
	"reflect"
	"time"
)

type Config struct {
	DB_Host string `env:"DB_HOST" default:"localhost"`
	DB_Port int    `env:"DB_PORT" default:"5432"`
	DB_User string `env:"DB_USER" default:""`
	DB_Pass string `env:"DB_PASS" default:""`
	DB_Name string `env:"DB_NAME" default:"postgres"`
}

var Database Config

var (
	RedisChatExpiry = time.Since(time.Now().AddDate(0, 0, 1)) // 1 day
)

const (
	ENV     = "env"
	DEFAULT = "default"
)

func setConfig(destination Config) {
	v := reflect.ValueOf(destination)
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get(ENV)
		defaultTag := v.Type().Field(i).Tag.Get(DEFAULT)

		if tag == "" || tag == "-" {
			continue
		}

		a := reflect.Indirect(reflect.ValueOf(destination))
		EnvVar, Info := loadFromEnv(tag, defaultTag)
		if Info != "" {
			fmt.Println("Missing environment configuration for '" + a.Type().Field(i).Name + "', Loading default setting!")
		}

		switch t := v.Type().Field(i).Type.Kind(); t {
		case reflect.String:
			fmt.Println("(string)")
			reflect.ValueOf(&destination).Elem().Field(i).SetString(EnvVar)
		case reflect.Int:
			fmt.Println("(int)")
			reflect.ValueOf(&destination).Elem().Field(i).SetInt(int64(utils.ToInt(EnvVar)))
		}
		fmt.Println(destination)
	}
	Database = destination
}

func loadFromEnv(tag string, defaultTag string) (string, string) {
	envVar := utils.GetEnv(tag)
	if envVar == "" && defaultTag != "" {
		envVar = defaultTag
		return envVar, "1"
	}
	return envVar, ""
}

func GetConfiguration() Config {
	return Database
}

func init() {
	setConfig(Database)

	fmt.Printf("Service configuration : %+v\n ", Database)
}
