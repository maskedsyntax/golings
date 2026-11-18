// json_02: Strict config with defaults
//
// Implement ParseConfig. See README.md.
//
// Reference: https://gobyexample.com/json
package config

// Config is a server configuration.
type Config struct {
	Host  string `json:"host"`
	Port  int    `json:"port"`
	Debug bool   `json:"debug"`
}

// ParseConfig decodes data, applying defaults and rejecting unknown keys.
func ParseConfig(data []byte) (Config, error) {
	// TODO
	return Config{}, nil
}
