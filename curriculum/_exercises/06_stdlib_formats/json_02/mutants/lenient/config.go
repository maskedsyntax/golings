// Mutant: json.Unmarshal ignores unknown keys.
package config

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	Host  string `json:"host"`
	Port  int    `json:"port"`
	Debug bool   `json:"debug"`
}

func ParseConfig(data []byte) (Config, error) {
	cfg := Config{Host: "localhost", Port: 8080}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}
	if cfg.Port < 1 || cfg.Port > 65535 {
		return Config{}, fmt.Errorf("port %d out of range", cfg.Port)
	}
	return cfg, nil
}
