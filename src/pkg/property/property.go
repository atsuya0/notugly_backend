package ctx

import (
	"context"
)

const (
	userId = "userId"
)

func SetUserId(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, userId, value)
}

func GetUserId(ctx context.Context) string {
	if value, ok := ctx.Value(userId).(string); ok {
		return value
	}
	return ""
}
