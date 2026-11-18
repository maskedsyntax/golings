package variables

import "testing"

func TestBalance(t *testing.T) {
	tests := []struct {
		name          string
		in            []int
		wantBalance   int
		wantOverdrawn bool
	}{
		{"recovers after dipping", []int{10, -20, 30}, 20, true},
		{"never negative", []int{5, 5, -10}, 0, false},
		{"ends negative", []int{-1}, -1, true},
		{"empty", nil, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, o := Balance(tt.in)
			if b != tt.wantBalance || o != tt.wantOverdrawn {
				t.Errorf("Balance(%v) = (%d, %v), want (%d, %v)", tt.in, b, o, tt.wantBalance, tt.wantOverdrawn)
			}
		})
	}
}
