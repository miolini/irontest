package server

import (
	"bytes"
	"io"
	"net/http"
)

func readBody(req *http.Request) (string, error) {
	buf := bytes.Buffer{}
	_, err := io.Copy(&buf, req.Body)
	return string(buf.Bytes()), err
}
