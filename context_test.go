package goer

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestContextCorrelationId(t *testing.T) {

	var (
		id = uuid.Must(uuid.NewV4()).String()
	)

	ctx := context.Background()

	ctx = SetCorrelationId(ctx, id)

	assert.Equal(t, id, GetCorrelationId(ctx))
}

func TestContextReqeustId(t *testing.T) {

	var (
		id = uuid.Must(uuid.NewV4()).String()
	)

	ctx := context.Background()

	ctx = SetRequestId(ctx, id)

	assert.Equal(t, id, GetRequestId(ctx))
}

func TestContextAppId(t *testing.T) {

	var (
		id = uuid.Must(uuid.NewV4()).String()
	)

	ctx := context.Background()

	ctx = SetAppId(ctx, id)

	assert.Equal(t, id, GetAppId(ctx))
}

func TestContextUserId(t *testing.T) {

	var (
		id = uuid.Must(uuid.NewV4()).String()
	)

	ctx := context.Background()

	ctx = SetUserId(ctx, id)

	assert.Equal(t, id, GetUserId(ctx))
}

func TestContextUserName(t *testing.T) {

	var (
		val = uuid.Must(uuid.NewV4()).String()
	)

	ctx := context.Background()

	ctx = SetUserName(ctx, val)

	assert.Equal(t, val, GetUserName(ctx))
}

func TestContextGetWhenNoExist(t *testing.T) {

	ctx := context.Background()

	id := GetCorrelationId(ctx)

	assert.Equal(t, "", id)
}
