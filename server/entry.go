package server

import (
	"sync"
)

type Entry struct {
	sync.Mutex
	value  string
	lockID string
}

func newEntry() *Entry {
	e := &Entry{}
	return e
}

func newEntryInterface() interface{} {
	return newEntry()
}
