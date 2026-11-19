package reqid

import (
	"context"
	"testing"
)

func TestRequestID(t *testing.T) {
	ctx := WithRequestID(context.Background(), "req-42")
	if id, ok := RequestID(ctx); !ok || id != "req-42" {
		t.Fatalf("RequestID = %q, %v; want \"req-42\", true", id, ok)
	}
	if id, ok := RequestID(context.Background()); ok || id != "" {
		t.Fatalf("RequestID(no ID) = %q, %v; want \"\", false", id, ok)
	}
}

func TestNoCollisions(t *testing.T) {
	ctx := context.WithValue(context.Background(), "requestID", "spoofed")
	if id, ok := RequestID(ctx); ok {
		t.Fatalf("RequestID found %q, set by someone else's string key", id)
	}
	ctx = WithRequestID(context.Background(), "real")
	ctx = context.WithValue(ctx, "requestID", "spoofed")
	if id, _ := RequestID(ctx); id != "real" {
		t.Fatalf("RequestID = %q after another package set the string key; want \"real\"", id)
	}
}

func TestForAudit(t *testing.T) {
	parent, cancel := context.WithCancel(WithRequestID(context.Background(), "req-7"))
	audit := ForAudit(parent)
	cancel()
	if parent.Err() == nil {
		t.Fatal("test setup: parent should be cancelled")
	}
	if err := audit.Err(); err != nil {
		t.Errorf("ForAudit context was cancelled with the request: %v", err)
	}
	if id, ok := RequestID(audit); !ok || id != "req-7" {
		t.Errorf("RequestID(ForAudit(ctx)) = %q, %v; want \"req-7\", true", id, ok)
	}
}
