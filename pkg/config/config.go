package config

import (
	"log"
	"os"
	"strconv"
)

type AppConfig struct {
	AppName                 string
	Port                    int
	DbHost                  string
	DbPort                  int
	DbUser                  string
	DbPassword              string
	DbName                  string
	OAuthGoogleClientID     string
	OAuthGoogleClientSecret string
	OAuthGoogleCallbackURL  string
	SessionSecret           string
	SessionDuration         int
	SessionKey              string
	StripeKey               string
	StripeSecret            string
}

var defaultConfig = AppConfig{
	AppName:         "application",
	Port:            8080,
	DbHost:          "localhost",
	DbPort:          5432,
	DbUser:          "admin",
	DbPassword:      "1234",
	DbName:          "testlocal",
	SessionDuration: 3600,
	SessionKey:      "session",
}

var appConfig AppConfig

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
			log.Println("There's an error parsing an environment variable:", v)
			return defaultValue
		}

		return parsedValue
	}

	return defaultValue
}

func SetConfig() {
	appConfig = AppConfig{
		AppName:                 withDefault("APP_NAME", defaultConfig.AppName),
		Port:                    withDefaultInt("APP_PORT", defaultConfig.Port),
		DbHost:                  withDefault("POSTGRES_HOSTNAME", defaultConfig.DbHost),
		DbPort:                  withDefaultInt("POSTGRES_PORT", defaultConfig.DbPort),
		DbUser:                  withDefault("POSTGRES_USER", defaultConfig.DbUser),
		DbPassword:              withDefault("POSTGRES_PASSWORD", defaultConfig.DbPassword),
		DbName:                  withDefault("POSTGRES_DB", defaultConfig.DbName),
		OAuthGoogleClientID:     withDefault("OAUTH_GOOGLE_CLIENT_ID", defaultConfig.OAuthGoogleClientID),
		OAuthGoogleClientSecret: withDefault("OAUTH_GOOGLE_CLIENT_SECRET", defaultConfig.OAuthGoogleClientSecret),
		OAuthGoogleCallbackURL:  withDefault("OAUTH_GOOGLE_CALLBACK_URL", defaultConfig.OAuthGoogleCallbackURL),
		SessionSecret:           withDefault("SESSION_SECRET", defaultConfig.SessionSecret),
		SessionDuration:         withDefaultInt("SESSION_DURATION", defaultConfig.SessionDuration),
		SessionKey:              withDefault("SESSION_KEY", defaultConfig.SessionSecret),
		StripeKey:               withDefault("STRIPE_KEY", defaultConfig.StripeKey),
		StripeSecret:            withDefault("STRIPE_SECRET", defaultConfig.StripeSecret),
	}
}

func GetConfig() AppConfig {
	return appConfig
}
