// Mutant: treats any decode error as the end of input.
package levels

import (
	"encoding/json"
	"io"
)

func CountByLevel(r io.Reader) (map[string]int, error) {
	counts := make(map[string]int)
	dec := json.NewDecoder(r)
	for {
		var entry struct {
			Level string `json:"level"`
		}
		if err := dec.Decode(&entry); err != nil {
			return counts, nil
		}
		if entry.Level == "" {
			entry.Level = "unknown"
		}
		counts[entry.Level]++
	}
}
