package context

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func Background() context.Context {
	return context.Background()
}

func TODO() context.Context {
	return context.TODO()
}

func WithCancel(parent context.Context) (ctx context.Context, cancel context.CancelFunc) {
	return context.WithCancel(parent)
}

func WithValue(parent context.Context, key interface{}, val interface{}) context.Context {
	return context.WithValue(parent, key, val)
}

func WithDealtime(parent context.Context, d time.Time) (ctx context.Context, cancel context.CancelFunc) {
	return context.WithDeadline(parent, d)
}

func WithTimeout(parent context.Context, timeout time.Duration) (ctx context.Context, cancel context.CancelFunc) {
	return context.WithTimeout(parent, timeout)
}

type MetadataContext interface {
	context.Context

	Metadata() Metadata

	Get(key string) (string, bool)
	Set(key string, value interface{})

	SetCorrelationId(id string)
	GetCorrelationId() (string, bool)
	SetMessageId(messageId string)
	GetMessageId() (string, bool)
	SetTimestamp(time time.Time)
	GetTimestamp() (time.Time, bool)
	SetAppId(appId string)
	GetAppId() (string, bool)
	SetUserId(userId string)
	GetUserId() (string, bool)
	SetUserName(userName string)
	GetUserName() (string, bool)
}

type metadataCtx struct {
	context.Context
}

func WithMetadata(
	parent context.Context,
) MetadataContext {

	return &metadataCtx{
		Context: context.WithValue(parent, METADATA_KEY, make(Metadata)),
	}
}

func (ctx *metadataCtx) Metadata() Metadata {
	return ctx.Value(METADATA_KEY).(Metadata)
}

func (ctx *metadataCtx) Get(key string) (string, bool) {
	value, exist := ctx.Metadata()[MetadataKey(key)]

	return value, exist
}

func (ctx *metadataCtx) Set(key string, value interface{}) {
	ctx.Metadata()[MetadataKey(key)] = fmt.Sprintf("%v", value)
}

func (ctx *metadataCtx) get(key MetadataKey) (string, bool) {
	value, exist := ctx.Metadata()[key]

	return value, exist
}

func (ctx *metadataCtx) set(key MetadataKey, value interface{}) {
	ctx.Metadata()[key] = fmt.Sprintf("%v", value)
}

const (
	METADATA_KEY                = MetadataKey("metadata")
	METADATA_KEY_CORRELATION_ID = MetadataKey("correlation-id")
	METADATA_KEY_MESSAGE_ID     = MetadataKey("message-id")
	METADATA_KEY_TIMESTAMP      = MetadataKey("timestamp")
	METADATA_KEY_APP_ID         = MetadataKey("app-id")
	METADATA_KEY_USER_ID        = MetadataKey("user-id")
	METADATA_KEY_USER_NAME      = MetadataKey("user-name")
)

func (ctx *metadataCtx) SetCorrelationId(id string) {
	ctx.set(METADATA_KEY_CORRELATION_ID, id)
}

func (ctx *metadataCtx) GetCorrelationId() (string, bool) {
	return ctx.get(METADATA_KEY_CORRELATION_ID)
}

func (ctx *metadataCtx) SetMessageId(id string) {
	ctx.set(METADATA_KEY_MESSAGE_ID, id)
}

func (ctx *metadataCtx) GetMessageId() (string, bool) {
	return ctx.get(METADATA_KEY_MESSAGE_ID)
}

func (ctx *metadataCtx) SetAppId(appId string) {
	ctx.set(METADATA_KEY_APP_ID, appId)
}

func (ctx *metadataCtx) GetAppId() (string, bool) {
	return ctx.get(METADATA_KEY_APP_ID)
}

func (ctx *metadataCtx) SetUserId(userId string) {
	ctx.set(METADATA_KEY_USER_ID, userId)
}

func (ctx *metadataCtx) GetUserId() (string, bool) {
	return ctx.get(METADATA_KEY_USER_ID)
}

func (ctx *metadataCtx) SetUserName(userName string) {
	ctx.set(METADATA_KEY_USER_NAME, userName)
}

func (ctx *metadataCtx) GetUserName() (string, bool) {
	return ctx.get(METADATA_KEY_USER_NAME)
}

func (ctx *metadataCtx) SetTimestamp(time time.Time) {
	ctx.set(METADATA_KEY_TIMESTAMP, fmt.Sprintf("%d", time.UnixMilli()))
}

func (ctx *metadataCtx) GetTimestamp() (time.Time, bool) {
	if value, ok := ctx.get(METADATA_KEY_TIMESTAMP); ok {
		if timestamp, err := strconv.ParseInt(value, 10, 64); err == nil {
			return time.UnixMilli(timestamp), true
		}
	}

	return time.Time{}, false

}
