// context_03: Request-scoped values (reference solution)
package reqid

import "context"

// An unexported key type: no other package can build this key.
type requestIDKey struct{}

// WithRequestID returns a copy of ctx carrying id.
func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, id)
}

// RequestID returns the request ID stored by WithRequestID, if any.
func RequestID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(requestIDKey{}).(string)
	return id, ok
}

// ForAudit returns a context for work that must outlive the request: it keeps
// ctx's values but is not cancelled when ctx is.
func ForAudit(ctx context.Context) context.Context {
	return context.WithoutCancel(ctx)
}
