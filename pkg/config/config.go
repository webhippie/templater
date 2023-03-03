package config

// Generate defines the generate configuration.
type Generate struct {
	Prefix string `mapstructure:"prefix"`
	Output string `mapstructure:"output"`
}

// Logs defines the level and color for log configuration.
type Logs struct {
	Level  string `mapstructure:"level"`
	Pretty bool   `mapstructure:"pretty"`
	Color  bool   `mapstructure:"color"`
}

// Config defines the general configuration.
type Config struct {
	Generate Generate `mapstructure:"generate"`
	Logs     Logs     `mapstructure:"log"`
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}
