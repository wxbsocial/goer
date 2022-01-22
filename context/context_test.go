package context

import (
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetValue(t *testing.T) {
	ctx := Background()

	assert.NotNil(t, ctx.Value(METADATA_KEY))
}

func TestGetTimestamp(t *testing.T) {

	ctx := Background()

	now := time.Now()
	ctx.SetTimestamp(now)

	ts, ok := ctx.GetTimestamp()
	assert.True(t, ok)

	assert.Equal(t, now.UnixMilli(), ts.UnixMilli())
}

func TestGetCorrelationId(t *testing.T) {

	ctx := Background()

	id := uuid.Must(uuid.NewV4()).String()

	ctx.SetCorrelationId(id)

	id2, ok := ctx.GetCorrelationId()
	assert.True(t, ok)
	assert.Equal(t, id, id2)

}

func TestGetMessageId(t *testing.T) {

	ctx := Background()

	id := uuid.Must(uuid.NewV4()).String()

	ctx.SetMessageId(id)

	id2, ok := ctx.GetMessageId()
	assert.True(t, ok)
	assert.Equal(t, id, id2)

}

func TestGetAppId(t *testing.T) {

	ctx := Background()

	id := uuid.Must(uuid.NewV4()).String()

	ctx.SetAppId(id)

	id2, ok := ctx.GetAppId()
	assert.True(t, ok)
	assert.Equal(t, id, id2)

}

func TestGetUserId(t *testing.T) {

	ctx := Background()

	id := uuid.Must(uuid.NewV4()).String()

	ctx.SetUserId(id)

	id2, ok := ctx.GetUserId()
	assert.True(t, ok)
	assert.Equal(t, id, id2)

}

func TestGetUserName(t *testing.T) {

	ctx := Background()

	id := uuid.Must(uuid.NewV4()).String()

	ctx.SetUserName(id)

	id2, ok := ctx.GetUserName()
	assert.True(t, ok)
	assert.Equal(t, id, id2)

}
