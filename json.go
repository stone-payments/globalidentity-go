package globalidentity

import (
	"bytes"
	"encoding/json"
	"io"
)

func ToJson(s interface{}) (io.Reader, error) {
	b, err := json.Marshal(s)
	return bytes.NewBuffer(b), err
}

func FromJson(s interface{}, r io.Reader) error {
	decoder := json.NewDecoder(r)
	decoder.UseNumber()
	return decoder.Decode(s)
}
