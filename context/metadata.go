package context

import (
	"encoding/json"
)

type MetadataKey string

type Metadata map[MetadataKey]string

func (metadata Metadata) Bytes() ([]byte, error) {

	return json.Marshal(metadata)
}

func ParseMetadata(bytes []byte) (Metadata, error) {

	var metadata map[MetadataKey]string

	if err := json.Unmarshal(bytes, &metadata); err != nil {

		return nil, err
	}

	return metadata, nil
}
