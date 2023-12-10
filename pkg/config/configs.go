package config

type AppConfig struct {
	Port int
}

var defaultConfig = AppConfig{
	Port: 8080,
}

func GetConfig() AppConfig {
	return defaultConfig
}
