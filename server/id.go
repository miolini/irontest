package server

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func generateID(lenght int) string {
	buf := make([]byte, lenght/2+1)
	n, err := rand.Read(buf)
	if err != nil {
		panic(fmt.Errorf("generate id error: %s", err))
	} else if n != len(buf) {
		panic(fmt.Errorf("len of date less than expected: %d != %d", len(buf), n))
	}
	return hex.EncodeToString(buf)[:lenght]
}
