package goer

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

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

type metadataContext struct {
	context.Context
}

func NewMetadataContext(
	parent context.Context,
) MetadataContext {

	return &metadataContext{
		Context: context.WithValue(parent, METADATA_KEY, make(Metadata)),
	}
}

func (ctx *metadataContext) Metadata() Metadata {
	return ctx.Value(METADATA_KEY).(Metadata)
}

func (ctx *metadataContext) Get(key string) (string, bool) {
	value, exist := ctx.Metadata()[MetadataKey(key)]

	return value, exist
}

func (ctx *metadataContext) Set(key string, value interface{}) {
	ctx.Metadata()[MetadataKey(key)] = fmt.Sprintf("%v", value)
}

func (ctx *metadataContext) get(key MetadataKey) (string, bool) {
	value, exist := ctx.Metadata()[key]

	return value, exist
}

func (ctx *metadataContext) set(key MetadataKey, value interface{}) {
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

func (ctx *metadataContext) SetCorrelationId(id string) {
	ctx.set(METADATA_KEY_CORRELATION_ID, id)
}

func (ctx *metadataContext) GetCorrelationId() (string, bool) {
	return ctx.get(METADATA_KEY_CORRELATION_ID)
}

func (ctx *metadataContext) SetMessageId(id string) {
	ctx.set(METADATA_KEY_MESSAGE_ID, id)
}

func (ctx *metadataContext) GetMessageId() (string, bool) {
	return ctx.get(METADATA_KEY_MESSAGE_ID)
}

func (ctx *metadataContext) SetAppId(appId string) {
	ctx.set(METADATA_KEY_APP_ID, appId)
}

func (ctx *metadataContext) GetAppId() (string, bool) {
	return ctx.get(METADATA_KEY_APP_ID)
}

func (ctx *metadataContext) SetUserId(userId string) {
	ctx.set(METADATA_KEY_USER_ID, userId)
}

func (ctx *metadataContext) GetUserId() (string, bool) {
	return ctx.get(METADATA_KEY_USER_ID)
}

func (ctx *metadataContext) SetUserName(userName string) {
	ctx.set(METADATA_KEY_USER_NAME, userName)
}

func (ctx *metadataContext) GetUserName() (string, bool) {
	return ctx.get(METADATA_KEY_USER_NAME)
}

func (ctx *metadataContext) SetTimestamp(time time.Time) {
	ctx.set(METADATA_KEY_TIMESTAMP, fmt.Sprintf("%d", time.UnixMilli()))
}

func (ctx *metadataContext) GetTimestamp() (time.Time, bool) {
	if value, ok := ctx.get(METADATA_KEY_TIMESTAMP); ok {
		if timestamp, err := strconv.ParseInt(value, 10, 64); err == nil {
			return time.UnixMilli(timestamp), true
		}
	}

	return time.Time{}, false

}
