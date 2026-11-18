// Mutant: closes explicitly, but not when the read fails.
package total

import "io"

type Opener func(name string) (io.ReadCloser, error)

func TotalSize(open Opener, names []string) (int, error) {
	total := 0
	for _, name := range names {
		f, err := open(name)
		if err != nil {
			return 0, err
		}
		b, err := io.ReadAll(f)
		if err != nil {
			return 0, err
		}
		f.Close()
		total += len(b)
	}
	return total, nil
}
