// json_02: Strict config with defaults (reference solution)
package config

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Config is a server configuration.
type Config struct {
	Host  string `json:"host"`
	Port  int    `json:"port"`
	Debug bool   `json:"debug"`
}

// ParseConfig decodes data, applying defaults and rejecting unknown keys.
func ParseConfig(data []byte) (Config, error) {
	cfg := Config{Host: "localhost", Port: 8080} // absent keys keep these
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("parsing config: %w", err)
	}
	if cfg.Port < 1 || cfg.Port > 65535 {
		return Config{}, fmt.Errorf("port %d out of range", cfg.Port)
	}
	return cfg, nil
}
