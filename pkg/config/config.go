package config

import "github.com/spf13/viper"

type Config struct {
	APPENV     string `mapstructure:"APP_ENV"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPass     string `mapstructure:"DB_PASS"`
	DBPort     int32  `mapstructure:"DB_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	CacheUser  string `mapstructure:"CACHE_USER"`
	CachePass  string `mapstructure:"CACHE_PASS"`
	CachePort  int32  `mapstructure:"CACHE_PORT"`
	CacheHost  string `mapstructure:"CACHE_HOST"`
	StreamUser string `mapstructure:"STREAM_USER"`
	StreamPass string `mapstructure:"STREAM_PASS"`
	StreamPort int32  `mapstructure:"STREAM_PORT"`
	StreamHost string `mapstructure:"STREAM_HOST"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile("../../.env")
	viper.SetDefault("APP_ENV", "development")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, err
		}
	}

	err = viper.Unmarshal(&config)
	return config, err
}
