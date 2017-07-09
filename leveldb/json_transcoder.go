package leveldb

import "encoding/json"

type transcoder struct {
}

func (t transcoder) Decode(bytes []byte, out interface{}) error {
	return json.Unmarshal(bytes, &out)
}

func (t transcoder) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}
