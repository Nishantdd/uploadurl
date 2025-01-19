package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Server   ServerConfig
	App      AppConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	AWS      AWSConfig
	JWT      JWTConfig
	OAuth    OAuthConfig
}

type ServerConfig struct {
	Address string
}

type AppConfig struct {
	Address string
}

type PostgresConfig struct {
	URI string
}

type RedisConfig struct {
	Address  string
	Password string
}

type AWSConfig struct {
	S3BucketName string
	Region       string
	AccessKey    string
	SecretKey    string
}

type JWTConfig struct {
	JWTSecret string
}

type OAuthConfig struct {
	OAuthSecret        string
	RedirectURL        string
	GoogleClientId     string
	GoogleClientSecret string
	SessionName        string
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Address: os.Getenv("SERVER_ADDRESS"),
		},
		App: AppConfig{
			Address: os.Getenv("APP_ADDRESS"),
		},
		Postgres: PostgresConfig{
			URI: os.Getenv("DATABASE_URL"),
		},
		Redis: RedisConfig{
			Address:  os.Getenv("REDIS_ADDRESS"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
		AWS: AWSConfig{
			S3BucketName: os.Getenv("S3_BUCKET"),
			Region:       os.Getenv("AWS_REGION"),
			AccessKey:    os.Getenv("AWS_ACCESS_KEY"),
			SecretKey:    os.Getenv("AWS_SECRET_KEY"),
		},
		JWT: JWTConfig{
			JWTSecret: os.Getenv("JWT_SECRET"),
		},
		OAuth: OAuthConfig{
			OAuthSecret:        os.Getenv("OAUTH_SECRET"),
			RedirectURL:        os.Getenv("REDIRECT_URL"),
			GoogleClientId:     os.Getenv("GOOGLE_CLIENT_ID"),
			GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			RedirectURL:        os.Getenv("REDIRECT_URL"),
		},
	}
}
