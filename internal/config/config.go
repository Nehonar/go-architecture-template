package config

type Config struct {
	// Example:
	// HTTPPort string
	// DatabaseURL string
}

func Load() *Config {
	// TODO: load values from env or config file in real projects
	return &Config{}
}
