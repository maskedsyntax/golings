// Mutant: ForAudit fixed, key still a string.
package reqid

import "context"

func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, "requestID", id)
}

func RequestID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value("requestID").(string)
	return id, ok
}

func ForAudit(ctx context.Context) context.Context {
	return context.WithoutCancel(ctx)
}
