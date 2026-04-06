package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort  string
	AppEnv   string
	DB       DBConfig
	Auth     AuthConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func (d DBConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		d.Host, d.Port, d.User, d.Password, d.Name)
}

type AuthConfig struct {
	JWTSecret     string
	AdminUsername string
	AdminPassword string
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	return &Config{
		AppPort: getEnv("APP_PORT", "8081"),
		AppEnv:  getEnv("APP_ENV", "development"),
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "127.0.0.1"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "wayt_page"),
		},
		Auth: AuthConfig{
			JWTSecret:     getEnv("JWT_SECRET", "changeme"),
			AdminUsername: getEnv("ADMIN_USERNAME", "admin"),
			AdminPassword: getEnv("ADMIN_PASSWORD", ""),
		},
	}, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
