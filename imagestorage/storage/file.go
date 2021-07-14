package storage

import (
	"io"
)

type FileStorage struct {
	basePath string
}

func (fs *FileStorage) Put(reader io.Reader, name string) int {
	return -1
}

func (fs *FileStorage) Get(id int, writer io.Writer) bool {
	return true
}

func (fs *FileStorage) GetAllIds() []int {
	return []int{-1}
}

func (fs *FileStorage) HasId(id int) bool {
	return contains(fs.GetAllIds(), id)
}
