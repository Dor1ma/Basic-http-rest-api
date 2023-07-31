package store

type Config struct {
	DataBaseURL string `json:"data_base_url"`
}

func NewConfig() *Config {
	return &Config{}
}
