// Mutant: key fixed, ForAudit returns ctx itself.
package reqid

import "context"

type key struct{}

func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, key{}, id)
}

func RequestID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(key{}).(string)
	return id, ok
}

func ForAudit(ctx context.Context) context.Context {
	return ctx
}
