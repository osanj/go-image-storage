package storage

type FileStorage struct {
	basePath string
}

func (fs *FileStorage) put(data []byte, name string) int {
	return -1
}

func (fs *FileStorage) get(id int) []byte {
	return []byte{0}
}

func (fs *FileStorage) getAllIds() []int {
	return []int{-1}
}

func (fs *FileStorage) hasId(id int) bool {
	return contains(fs.getAllIds(), id)
}
