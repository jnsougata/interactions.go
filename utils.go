package main

import (
	"bytes"
	"encoding/json"
	"io"
)

func ReaderFromMap(v any) io.Reader {
	b, _ := json.Marshal(v)
	return bytes.NewReader(b)
}
