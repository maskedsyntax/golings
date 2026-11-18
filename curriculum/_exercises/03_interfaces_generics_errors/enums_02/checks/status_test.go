package status

import "testing"

func TestString(t *testing.T) {
	for _, tt := range []struct {
		s    Status
		want string
	}{
		{Pending, "pending"}, {Active, "active"}, {Suspended, "suspended"}, {Closed, "closed"},
		{Status(4), "Status(4)"}, {Status(-1), "Status(-1)"},
	} {
		if got := tt.s.String(); got != tt.want {
			t.Errorf("Status(%d).String() = %q, want %q", int(tt.s), got, tt.want)
		}
	}
}

func TestParseRoundTrip(t *testing.T) {
	for _, s := range []Status{Pending, Active, Suspended, Closed} {
		got, ok := ParseStatus(s.String())
		if !ok || got != s {
			t.Errorf("ParseStatus(%q) = %d, %v, want %d, true", s.String(), int(got), ok, int(s))
		}
	}
	for _, bad := range []string{"", "Active", "Status(4)", "deleted"} {
		if _, ok := ParseStatus(bad); ok {
			t.Errorf("ParseStatus(%q) succeeded, want false", bad)
		}
	}
}
