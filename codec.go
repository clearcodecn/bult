package blut

import "encoding/json"

type Codec interface {
	Marshal(v interface{}) ([]byte, error)
	UnMarshal(data []byte, v interface{}) error
}

type jsonCodec struct{}

func (j jsonCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (j jsonCodec) UnMarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

var (
	defaultCodec = &jsonCodec{}
)
