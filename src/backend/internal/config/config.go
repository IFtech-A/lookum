package config

import "fmt"

//Config holds configuration parameters for server
type Config struct {
	BindAddr        string `json:"bind_addr"`
	PortHTTP        int    `json:"port_http" env:"PORT_HTTP"`
	PortHTTPS       int    `json:"port_https" env:"PORT_HTTPS"`
	LogLevel        string `json:"log_level"`
	DatabaseURL     string `json:"database_url"`
	DbHost          string `json:"db_host" env:"DBHOST"`
	DbPort          int    `json:"db_port" env:"DBPORT"`
	DbUser          string `json:"db_user" env:"DBUSER"`
	DbPassword      string `json:"db_pass" env:"DBPASS"`
	DbName          string `json:"db_name" env:"DBNAME"`
	DbSSLMode       string `json:"db_ssl_mode" env:"DBSSLMODE"`
	SessionKey      string `env:"SESSION_KEY"`
	PrivateKeyPath  string `json:"pkey_path"`
	CertificatePath string `json:"cert_path"`
}

//NewConfig returns default configuration structure
func NewConfig() *Config {
	return &Config{
		BindAddr:   "0.0.0.0",
		PortHTTP:   80,
		PortHTTPS:  443,
		LogLevel:   "debug",
		DbHost:     "postgres",
		DbPort:     5432,
		DbUser:     "dbuser",
		DbPassword: "1234qwer!",
		DbName:     "lookum",
		DbSSLMode:  "disable",
		SessionKey: "secret",
	}
}

func (c *Config) GetDatabaseURL() string {
	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		c.DbHost,
		c.DbPort,
		c.DbUser,
		c.DbPassword,
		c.DbName,
		c.DbSSLMode,
	)
}
