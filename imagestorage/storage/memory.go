package storage

import (
	"bytes"
	"io"
	"log"
	"sync"
)

type MemoryStorageElement struct {
	metadata StorageItemMetadata
	bytes    []byte
}

type MemoryStorage struct {
	data   map[int]*MemoryStorageElement
	nextId int
	mutex  sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{data: map[int]*MemoryStorageElement{}, nextId: 1, mutex: sync.Mutex{}}
}

func (ms *MemoryStorage) Put(reader io.Reader, metadata StorageItemMetadata) int {
	ms.mutex.Lock()
	id := ms.nextId
	ms.nextId++
	ms.mutex.Unlock()

	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(reader)
	if err != nil {
		log.Fatal(err)
		return -1
	}
	ms.data[id] = &MemoryStorageElement{metadata: metadata, bytes: buffer.Bytes()}
	return id
}

func (ms *MemoryStorage) GetBytes(id int, writer io.Writer) bool {
	element, found := ms.data[id]
	if !found {
		return false
	}
	_, err := writer.Write(element.bytes)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (ms *MemoryStorage) GetMetadata(id int) *StorageItemMetadata {
	element, found := ms.data[id]
	if !found {
		return nil
	}
	return &element.metadata
}

func (ms *MemoryStorage) GetAllIds() []int {
	keys := make([]int, 0, len(ms.data))
	for k := range ms.data {
		keys = append(keys, k)
	}
	return keys
}

func (ms *MemoryStorage) HasId(id int) bool {
	return contains(ms.GetAllIds(), id)
}
