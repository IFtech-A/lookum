package apiserver

//Config holds configuration parameters for the apiserver
type Config struct {
	BindAddr        string `json:"bind_addr"`
	LogLevel        string `json:"log_level"`
	DatabaseURL     string `json:"database_url"`
	DbHost          string `json:"db_host" env:"DBHOST"`
	DbPort          int    `json:"db_port" env:"DBPORT"`
	DbUser          string `json:"db_user" env:"DBUSER"`
	DbPassword      string `json:"db_pass" env:"DBPASS"`
	DbName          string `json:"db_name" env:"DBNAME"`
	DbSSLMode       string `json:"db_ssl_mode" env:"DBSSLMODE"`
	PrivateKeyPath  string `json:"pkey_path"`
	CertificatePath string `json:"cert_path"`
}

//NewConfig returns default configuration structure
func NewConfig() *Config {
	return &Config{
		BindAddr:   "0.0.0.0",
		LogLevel:   "debug",
		DbHost:     "postgres",
		DbPort:     5432,
		DbUser:     "dbuser",
		DbPassword: "1234qwer!",
		DbName:     "lookum",
		DbSSLMode:  "disable",
	}
}
