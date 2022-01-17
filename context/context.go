package context

import (
	"context"
	"fmt"
	"strconv"
	"time"

	md "github.com/wxbsocial/goer/metadata"
)

func Background() Context {
	return newCtx(context.Background())
}

func TODO() Context {
	return newCtx(context.TODO())
}

func WithCancel(parent context.Context) (ctx Context, cancel context.CancelFunc) {
	ctx0, cancel0 := context.WithCancel(parent)

	return newCtx(ctx0), cancel0
}

func WithValue(parent context.Context, key interface{}, val interface{}) context.Context {
	return newCtx(context.WithValue(parent, key, val))

}

func WithDealtime(parent context.Context, d time.Time) (ctx context.Context, cancel context.CancelFunc) {
	ctx0, cancel0 := context.WithDeadline(parent, d)
	return newCtx(ctx0), cancel0
}

func WithTimeout(parent context.Context, timeout time.Duration) (ctx context.Context, cancel context.CancelFunc) {
	ctx0, cancel0 := context.WithTimeout(parent, timeout)
	return newCtx(ctx0), cancel0
}

func WithMetadata(parent context.Context) Context {
	return newCtx(parent)
}

type Context interface {
	context.Context

	Metadata() md.Metadata

	Get(key md.MetadataKey) (string, bool)
	Set(key md.MetadataKey, value string)

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

type ctx struct {
	context.Context
}

func newCtx(
	parent context.Context,
) Context {

	return &ctx{
		Context: context.WithValue(parent, METADATA_KEY, make(md.Metadata)),
	}
}

func (ctx *ctx) Metadata() md.Metadata {
	return ctx.Value(METADATA_KEY).(md.Metadata)
}

func (ctx *ctx) Get(key md.MetadataKey) (string, bool) {
	value, exist := ctx.Metadata()[key]

	return value, exist
}

func (ctx *ctx) Set(key md.MetadataKey, value string) {
	ctx.Metadata()[key] = fmt.Sprintf("%v", value)
}

const (
	METADATA_KEY                = md.MetadataKey("metadata")
	METADATA_KEY_CORRELATION_ID = md.MetadataKey("correlation-id")
	METADATA_KEY_MESSAGE_ID     = md.MetadataKey("message-id")
	METADATA_KEY_TIMESTAMP      = md.MetadataKey("timestamp")
	METADATA_KEY_APP_ID         = md.MetadataKey("app-id")
	METADATA_KEY_USER_ID        = md.MetadataKey("user-id")
	METADATA_KEY_USER_NAME      = md.MetadataKey("user-name")
)

func (ctx *ctx) SetCorrelationId(id string) {
	ctx.Set(METADATA_KEY_CORRELATION_ID, id)
}

func (ctx *ctx) GetCorrelationId() (string, bool) {
	return ctx.Get(METADATA_KEY_CORRELATION_ID)
}

func (ctx *ctx) SetMessageId(id string) {
	ctx.Set(METADATA_KEY_MESSAGE_ID, id)
}

func (ctx *ctx) GetMessageId() (string, bool) {
	return ctx.Get(METADATA_KEY_MESSAGE_ID)
}

func (ctx *ctx) SetAppId(appId string) {
	ctx.Set(METADATA_KEY_APP_ID, appId)
}

func (ctx *ctx) GetAppId() (string, bool) {
	return ctx.Get(METADATA_KEY_APP_ID)
}

func (ctx *ctx) SetUserId(userId string) {
	ctx.Set(METADATA_KEY_USER_ID, userId)
}

func (ctx *ctx) GetUserId() (string, bool) {
	return ctx.Get(METADATA_KEY_USER_ID)
}

func (ctx *ctx) SetUserName(userName string) {
	ctx.Set(METADATA_KEY_USER_NAME, userName)
}

func (ctx *ctx) GetUserName() (string, bool) {
	return ctx.Get(METADATA_KEY_USER_NAME)
}

func (ctx *ctx) SetTimestamp(time time.Time) {
	ctx.Set(METADATA_KEY_TIMESTAMP, fmt.Sprintf("%d", time.UnixMilli()))
}

func (ctx *ctx) GetTimestamp() (time.Time, bool) {
	if value, ok := ctx.Get(METADATA_KEY_TIMESTAMP); ok {
		if timestamp, err := strconv.ParseInt(value, 10, 64); err == nil {
			return time.UnixMilli(timestamp), true
		}
	}

	return time.Time{}, false

}
