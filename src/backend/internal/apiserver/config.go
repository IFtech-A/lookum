package apiserver

//Config holds configuration parameters for the apiserver
type Config struct {
	BindAddr        string `json:"bind_addr"`
	LogLevel        string `json:"log_level"`
	DatabaseURL     string `json:"database_url"`
	PrivateKeyPath  string `json:"pkey_path"`
	CertificatePath string `json:"cert_path"`
}

//NewConfig returns default configuration structure
func NewConfig() *Config {
	return &Config{
		BindAddr:    "0.0.0.0",
		LogLevel:    "debug",
		DatabaseURL: "host=192.168.1.104 port=5432 user=postgres password=1234qwer! dbname=lookum sslmode=disable",
	}
}
