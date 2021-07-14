package storage

import (
	"io"
)

type StorageItemMetadata struct {
	Name     string
	MimeType string
}

type Storage interface {
	Put(reader io.Reader, metadata StorageItemMetadata) int
	GetBytes(id int, writer io.Writer) bool
	GetMetadata(id int) *StorageItemMetadata
	GetAllIds() []int
	HasId(id int) bool
}
