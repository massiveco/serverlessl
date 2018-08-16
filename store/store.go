package store

import (
	"bytes"
)

// Store interface for fetching files from a object store
type Store interface {
	FetchFile(filename string, buf *bytes.Buffer) error
	PutFile(filename string, buf *bytes.Reader) error
	PutPublicFile(filename string, buf *bytes.Reader) error
}
