package main

import "fmt"

// PostgresConfig represents the configuration
// for a PostgreSQL database.
type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// DefaultPostgresConfig is used to create a new PostgresConfig.
func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "ec2-54-157-88-70.compute-1.amazonaws.com",
		Port:     5432,
		User:     "czbyfpbnovxxsj",
		Password: "3ac3643181cc898fa2f80f3b9f77d4cdeb514062a5e3780d30493eaa7f9cb193",
		Name:     "d2jd17qkkc675u",
	}
}

// Dialect returns the specific dialect of SQL.
func (c PostgresConfig) Dialect() string {
	return "postgres"
}

// ConnectionInfo returns a string with the connection
// information formatted properly.
func (c PostgresConfig) ConnectionInfo() string {
	if c.Password == "" {
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s", c.Host, c.Port, c.User, c.Name)
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", c.Host, c.Port, c.User, c.Password, c.Name)
}

// Config represents the configuration
// for the application.
type Config struct {
	Port    int    `json:"port"`
	Env     string `json:"env"`
	Pepper  string `json:"pepper"`
	HMACKey string `json:"hmac_key"`
}

// DefaultConfig is used to create a new Config.
func DefaultConfig() Config {
	return Config{
		Port:    3000,
		Env:     "dev",
		Pepper:  "secret-random-key",
		HMACKey: "secret-hmac-key",
	}
}

// IsProd returns whether or not the
// application is in production.
func (c Config) IsProd() bool {
	return c.Env == "prod"
}
