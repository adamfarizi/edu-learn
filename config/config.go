// Halaman untuk menyambungkan koneksi dengan database

package config

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	Driver   string
}

type APIConfig struct {
	ApiPort string
}

type TokenConfig struct {
	ApplicationName     string
	JWTSignatureKey     []byte
	JWTSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

type Config struct {
	DBConfig
	APIConfig
	TokenConfig
}

func (c *Config) readConfig() error {
	err := godotenv.Load() // Memanggil .env
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	c.DBConfig = DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	c.APIConfig = APIConfig{
		ApiPort: "8080",
	}

	c.TokenConfig = TokenConfig{
		ApplicationName:     os.Getenv("APPLICATION_NAME"),   
		JWTSignatureKey:     []byte(os.Getenv("JWT_SECRET")), 
		JWTSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifeTime: time.Duration(1) * time.Hour, // 1 jam
	}

	if c.Host == "" || c.Port == "" || c.Username == "" || c.Password == "" || c.ApiPort == "" {
		return fmt.Errorf("required config")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cfg.readConfig()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
