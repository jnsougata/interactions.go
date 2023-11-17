package main

import (
	"bytes"
	"encoding/json"
	"io"
)

func ReaderFromAny(v any) io.Reader {
	b, _ := json.Marshal(v)
	return bytes.NewReader(b)
}
