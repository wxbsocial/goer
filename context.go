package goer

import "context"

type ContextKey string

var (
	X_REQUEST_ID     = ContextKey("X-Request-Id")
	X_CORRELATION_ID = ContextKey("X-Correlation-Id")
	X_APP_ID         = ContextKey("X-App-Id")
	X_USER_ID        = ContextKey("X-User-Id")
	X_USER_NAME      = ContextKey("X-User-Name")
)

func SetRequestId(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, X_REQUEST_ID, value)
}

func GetRequestId(ctx context.Context) string {
	if val, ok := ctx.Value(X_REQUEST_ID).(string); ok {
		return val
	}
	return ""

}

func SetCorrelationId(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, X_CORRELATION_ID, value)
}

func GetCorrelationId(ctx context.Context) string {
	if val, ok := ctx.Value(X_CORRELATION_ID).(string); ok {
		return val
	}
	return ""
}

func SetAppId(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, X_APP_ID, value)
}

func GetAppId(ctx context.Context) string {
	if val, ok := ctx.Value(X_APP_ID).(string); ok {
		return val
	}
	return ""
}

func SetUserId(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, X_USER_ID, value)
}

func GetUserId(ctx context.Context) string {
	if val, ok := ctx.Value(X_USER_ID).(string); ok {
		return val
	}
	return ""
}

func SetUserName(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, X_USER_NAME, value)
}

func GetUserName(ctx context.Context) string {
	if val, ok := ctx.Value(X_USER_NAME).(string); ok {
		return val
	}
	return ""
}
