package config

import (
	"log"
	"os"
	"strconv"
)

type AppConfig struct {
	Port       int
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string
}

var defaultConfig = AppConfig{
	Port:       8080,
	DbHost:     "localhost",
	DbPort:     5432,
	DbUser:     "admin",
	DbPassword: "1234",
	DbName:     "testlocal",
}

func withDefault(v, defaultValue string) string {
	value, found := os.LookupEnv(v)
	if found {
		return value
	}

	return defaultValue
}

func withDefaultInt(v string, defaultValue int) int {
	value, found := os.LookupEnv(v)
	if found {
		parsedValue, err := strconv.Atoi(value)
		if err != nil {
			log.Println("There's an error parsing an environment variable", v)
			return defaultValue
		}

		return parsedValue
	}

	return defaultValue
}

func GetConfig() AppConfig {
	appConfig := AppConfig{
		Port:       withDefaultInt("APP_PORT", defaultConfig.Port),
		DbHost:     withDefault("POSTGRES_HOSTNAME", defaultConfig.DbHost),
		DbPort:     withDefaultInt("POSTGRES_PORT", defaultConfig.DbPort),
		DbUser:     withDefault("POSTGRES_USER", defaultConfig.DbUser),
		DbPassword: withDefault("POSTGRES_PASSWORD", defaultConfig.DbPassword),
		DbName:     withDefault("POSTGRES_DB", defaultConfig.DbName),
	}

	return appConfig
}
