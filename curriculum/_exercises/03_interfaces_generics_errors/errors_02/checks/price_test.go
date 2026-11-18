package price

import (
	"errors"
	"strings"
	"testing"
)

var prices = map[string]int{"apple": 120, "free sample": 0, "broken": -5}

func TestPriceFound(t *testing.T) {
	for key, want := range map[string]int{"apple": 120, "free sample": 0} {
		if got, err := Price(prices, key); err != nil || got != want {
			t.Errorf("Price(%q) = %d, %v, want %d, nil", key, got, err, want)
		}
	}
}

func TestPriceNotFound(t *testing.T) {
	got, err := Price(prices, "kiwi")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("Price(\"kiwi\") error = %v; errors.Is(err, ErrNotFound) should be true", err)
	}
	if err == ErrNotFound {
		t.Error("the error should add context, not be ErrNotFound itself")
	}
	if !strings.Contains(err.Error(), "kiwi") || got != 0 {
		t.Errorf("Price(\"kiwi\") = %d, %q; want 0 and a message mentioning kiwi", got, err)
	}
	if errors.Is(err, ErrInvalid) {
		t.Error("a missing key should not be ErrInvalid")
	}
}

func TestPriceInvalid(t *testing.T) {
	got, err := Price(prices, "broken")
	if !errors.Is(err, ErrInvalid) || got != 0 || !strings.Contains(err.Error(), "broken") {
		t.Errorf("Price(\"broken\") = %d, %v; want 0 and an error wrapping ErrInvalid that mentions broken", got, err)
	}
	if errors.Is(err, ErrNotFound) {
		t.Error("a negative price should not be ErrNotFound")
	}
}
