// json_03: Streaming JSON lines (reference solution)
package levels

import (
	"encoding/json"
	"errors"
	"io"
)

// CountByLevel counts JSON log entries in r by their "level" field.
func CountByLevel(r io.Reader) (map[string]int, error) {
	counts := make(map[string]int)
	dec := json.NewDecoder(r)
	for {
		var entry struct {
			Level string `json:"level"`
		}
		err := dec.Decode(&entry)
		if errors.Is(err, io.EOF) {
			return counts, nil
		}
		if err != nil {
			return nil, err
		}
		if entry.Level == "" {
			entry.Level = "unknown"
		}
		counts[entry.Level]++
	}
}
