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
