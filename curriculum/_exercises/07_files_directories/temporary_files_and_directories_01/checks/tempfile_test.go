package tempfile

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWithTempFile(t *testing.T) {
	var seen string
	err := WithTempFile([]byte("hello"), func(p string) error {
		seen = p
		data, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		if string(data) != "hello" {
			t.Errorf("temp file contains %q, want \"hello\"", data)
		}
		// fn may delete or rename the file itself on every OS only if it's closed.
		return os.Rename(p, p+".moved")
	})
	if err != nil {
		t.Fatalf("WithTempFile returned an error: %v", err)
	}
	os.Remove(seen + ".moved")
	if base := filepath.Base(seen); !strings.HasPrefix(base, "golings-") || !strings.HasSuffix(base, ".txt") {
		t.Errorf("temp file name %q should look like golings-*.txt", base)
	}
}

func TestWithTempFileRemoves(t *testing.T) {
	var seen string
	if err := WithTempFile([]byte("x"), func(p string) error { seen = p; return nil }); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(seen); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("%s still exists after WithTempFile returned", seen)
	}
}

func TestWithTempFileRemovesOnError(t *testing.T) {
	boom := errors.New("boom")
	var seen string
	err := WithTempFile([]byte("x"), func(p string) error { seen = p; return boom })
	if !errors.Is(err, boom) {
		t.Fatalf("WithTempFile error = %v, want fn's error", err)
	}
	if _, err := os.Stat(seen); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("%s still exists after fn failed", seen)
	}
}
