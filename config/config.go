package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Debug      string `mapstructure:"DEBUG"`
	ServerHost string `mapstructure:"APP_SERVER_HOST"`
	ServerPort string `mapstructure:"APP_SERVER_PORT"`

	DBConn string `mapstructure:"DB_CONNECTION"`
	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USERNAME"`
	DBPass string `mapstructure:"DB_PASSWORD"`
	DBName string `mapstructure:"DB_NAME"`

	AWSRegion      string `mapstructure:"AWS_REGION"`
	AWSAccessKeyId string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWSSecretKey   string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	BucketName     string `mapstructure:"BUCKET_NAME"`

	GomailEmail    string `mapstructure:"GOMAIL_EMAIL"`
	GomailPassword string `mapstructure:"GOMAIL_PASSWORD"`
}

type JWTSecret struct {
	Secret string `mapstructure:"JWT_SECRET"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&config)
	return config, err
}

func LoadJWTSecret(path string) (secret JWTSecret, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return JWTSecret{}, err
	}

	err = viper.Unmarshal(&secret)
	return secret, err
}
