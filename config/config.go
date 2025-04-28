// Halaman untuk menyambungkan koneksi dengan database

package config

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	c.DBConfig = DBConfig{
		Host:     "localhost",
		Port:     "5432",
		Database: "edulearn_db", // Nama Database
		Username: "postgres",
		Password: "04042001",
		Driver:   "postgres",
	}

	c.APIConfig = APIConfig{
		ApiPort: "8080",
	}

	c.TokenConfig = TokenConfig{
		ApplicationName: "Edu Learn", // Nama Aplikasi
		JWTSignatureKey: []byte("Edu Learn JWT"), // Signature Key JWT
		JWTSigningMethod: jwt.SigningMethodHS256,
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
