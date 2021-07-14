package storage

type MemoryStorageElement struct {
	name string
	data []byte
}

type MemoryStorage struct {
	data []MemoryStorageElement
}

func (fs *MemoryStorage) put(data []byte, name string) int {
	return -1
}

func (fs *MemoryStorage) get(id int) []byte {
	return []byte{0}
}

func (fs *MemoryStorage) getAllIds() []int {
	return []int{-1}
}

func (fs *MemoryStorage) hasId(id int) bool {
	return contains(fs.getAllIds(), id)
}
