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

	ts := ctx.GetTimestamp()

	assert.Equal(t, now.UnixMilli(), ts.UnixMilli())
}

func TestGetTimestampWhenNotSet(t *testing.T) {

	ctx := Background()

	ts := ctx.GetTimestamp()

	assert.Equal(t, time.Time{}.UnixMilli(), ts.UnixMilli())
}

func TestGetCorrelationId(t *testing.T) {

	ctx := Background()

	id := uuid.Must(uuid.NewV4()).String()

	ctx.SetCorrelationId(id)

	id2 := ctx.GetCorrelationId()

	assert.Equal(t, id, id2)

}

func TestGetMessageId(t *testing.T) {

	ctx := Background()

	id := uuid.Must(uuid.NewV4()).String()

	ctx.SetMessageId(id)

	id2 := ctx.GetMessageId()

	assert.Equal(t, id, id2)

}

func TestGetAppId(t *testing.T) {

	ctx := Background()

	id := uuid.Must(uuid.NewV4()).String()

	ctx.SetAppId(id)

	id2 := ctx.GetAppId()

	assert.Equal(t, id, id2)

}

func TestGetUserId(t *testing.T) {

	ctx := Background()

	id := uuid.Must(uuid.NewV4()).String()

	ctx.SetUserId(id)

	id2 := ctx.GetUserId()

	assert.Equal(t, id, id2)

}

func TestGetUserName(t *testing.T) {

	ctx := Background()

	id := uuid.Must(uuid.NewV4()).String()

	ctx.SetUserName(id)

	id2 := ctx.GetUserName()

	assert.Equal(t, id, id2)

}

func TestMetadataTransfer(t *testing.T) {
	const (
		msgId = "001"
	)
	ctx := Background()
	ctx.SetMessageId(msgId)
	assert.Equal(t, msgId, ctx.GetMessageId())

	ctx2, _ := WithCancel(ctx)
	assert.Equal(t, msgId, ctx2.GetMessageId())

}

func TestCancelContext(t *testing.T) {

	ctx, cancel := WithCancel(Background())
	ctx2 := WithValue(ctx, METADATA_KEY_APP_ID, "001")

	select {
	case <-ctx2.Done():
		assert.True(t, false)
	default:
	}

	cancel()

	select {
	case <-ctx2.Done():
		assert.True(t, true)
	default:
	}

}
