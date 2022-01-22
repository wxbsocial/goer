package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindString(t *testing.T) {

	metadata := Metadata{}

	key := MetadataKey("message")
	val := "001"

	metadata[key] = val

	bytes, err := metadata.Bytes()
	assert.NoError(t, err)
	metadata2, err := ParseMetadata(bytes)
	assert.NoError(t, err)

	val2, ok := metadata2[key]
	assert.True(t, ok)
	assert.True(t, val == val2)

}

func TestUnion(t *testing.T) {
	md1 := Metadata{"A": "1"}
	md2 := Metadata{"B": "2"}
	var exist bool
	var val string
	assert.Equal(t, 1, len(md1))
	assert.Equal(t, 1, len(md2))
	val, exist = md1["A"]
	assert.True(t, exist)
	assert.Equal(t, val, "1")
	_, exist = md1["B"]
	assert.False(t, exist)
	md1.Union(md2)
	assert.Equal(t, 2, len(md1))
	assert.Equal(t, 1, len(md2))
	val, exist = md1["A"]
	assert.True(t, exist)
	assert.Equal(t, val, "1")
	val, exist = md1["B"]
	assert.Equal(t, val, "2")
	assert.True(t, exist)

}
