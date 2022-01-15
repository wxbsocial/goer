package context

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type Context interface {
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

type ctx struct {
	context.Context
}

func NewMetadataContext(
	parent context.Context,
) Context {

	return &ctx{
		Context: context.WithValue(parent, METADATA_KEY, make(Metadata)),
	}
}

func (ctx *ctx) Metadata() Metadata {
	return ctx.Value(METADATA_KEY).(Metadata)
}

func (ctx *ctx) Get(key string) (string, bool) {
	value, exist := ctx.Metadata()[MetadataKey(key)]

	return value, exist
}

func (ctx *ctx) Set(key string, value interface{}) {
	ctx.Metadata()[MetadataKey(key)] = fmt.Sprintf("%v", value)
}

func (ctx *ctx) get(key MetadataKey) (string, bool) {
	value, exist := ctx.Metadata()[key]

	return value, exist
}

func (ctx *ctx) set(key MetadataKey, value interface{}) {
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

func (ctx *ctx) SetCorrelationId(id string) {
	ctx.set(METADATA_KEY_CORRELATION_ID, id)
}

func (ctx *ctx) GetCorrelationId() (string, bool) {
	return ctx.get(METADATA_KEY_CORRELATION_ID)
}

func (ctx *ctx) SetMessageId(id string) {
	ctx.set(METADATA_KEY_MESSAGE_ID, id)
}

func (ctx *ctx) GetMessageId() (string, bool) {
	return ctx.get(METADATA_KEY_MESSAGE_ID)
}

func (ctx *ctx) SetAppId(appId string) {
	ctx.set(METADATA_KEY_APP_ID, appId)
}

func (ctx *ctx) GetAppId() (string, bool) {
	return ctx.get(METADATA_KEY_APP_ID)
}

func (ctx *ctx) SetUserId(userId string) {
	ctx.set(METADATA_KEY_USER_ID, userId)
}

func (ctx *ctx) GetUserId() (string, bool) {
	return ctx.get(METADATA_KEY_USER_ID)
}

func (ctx *ctx) SetUserName(userName string) {
	ctx.set(METADATA_KEY_USER_NAME, userName)
}

func (ctx *ctx) GetUserName() (string, bool) {
	return ctx.get(METADATA_KEY_USER_NAME)
}

func (ctx *ctx) SetTimestamp(time time.Time) {
	ctx.set(METADATA_KEY_TIMESTAMP, fmt.Sprintf("%d", time.UnixMilli()))
}

func (ctx *ctx) GetTimestamp() (time.Time, bool) {
	if value, ok := ctx.get(METADATA_KEY_TIMESTAMP); ok {
		if timestamp, err := strconv.ParseInt(value, 10, 64); err == nil {
			return time.UnixMilli(timestamp), true
		}
	}

	return time.Time{}, false

}
