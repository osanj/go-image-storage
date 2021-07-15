package storage

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type FileStorage struct {
	basePath          string
	createIfNotExists bool
	data              map[int]*StorageItemMetadata
	nextId            int
	mutex             sync.Mutex
}

func NewFileStorage(basePath string, createIfNotExists bool) *FileStorage {
	items := map[int]*StorageItemMetadata{}
	maxId := 0

	if createIfNotExists {
		err := os.MkdirAll(basePath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fname := f.Name()
		fpath := filepath.Join(basePath, fname)

		fi, err := os.Stat(fpath)
		if err != nil {
			log.Fatal(err)
		}

		if !fi.IsDir() {
			continue
		}

		dirId, err := strconv.Atoi(fname)
		if err != nil || dirId < 0 {
			log.Printf("skipping directory %s (name is not an id)", fname)
			continue
		}

		if dirId > maxId {
			maxId = dirId
		}

		files2, err := ioutil.ReadDir(fpath)
		if err != nil {
			log.Fatal(err)
		}

		if len(files2) != 1 {
			log.Printf("skipping directory %s (does not contain only 1 file)", fname)
			continue
		}

		f2 := files2[0]
		_, f2Name := filepath.Split(f2.Name())
		mimeTypeDefault := ""
		mimeType := mimeTypeDefault

		switch strings.ToLower(filepath.Ext(f2.Name())) {
		case ".jpeg":
			mimeType = "image/jpeg"
		case ".png":
			mimeType = "image/png"
		}

		if mimeType == mimeTypeDefault {
			log.Printf("skipping file %s (mimeType can not be inferred)", f2.Name())
			continue
		}

		items[dirId] = &StorageItemMetadata{Name: f2Name, MimeType: mimeType}
	}

	log.Printf("found %d images", len(items))
	return &FileStorage{basePath: basePath, data: items, nextId: maxId + 1, mutex: sync.Mutex{}}
}

func (fs *FileStorage) Put(reader io.Reader, metadata StorageItemMetadata) int {
	fs.mutex.Lock()
	id := fs.nextId
	fs.nextId++
	fs.mutex.Unlock()

	pathDir := filepath.Join(fs.basePath, strconv.Itoa(id))
	path := filepath.Join(pathDir, metadata.Name)

	err := os.MkdirAll(pathDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return -1
	}
	f, err2 := os.Create(path)
	if err2 != nil {
		log.Fatal(err2)
		return -1
	}
	_, err3 := io.Copy(f, reader)
	defer f.Close()
	if err3 != nil {
		log.Fatal(err3)
		return -1
	}

	fs.data[id] = &metadata
	return id
}

func (fs *FileStorage) GetBytes(id int, writer io.Writer) bool {
	metadata := fs.GetMetadata(id)
	if metadata == nil {
		return false
	}

	path := filepath.Join(fs.basePath, strconv.Itoa(id), metadata.Name)

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return false
	}
	_, err2 := io.Copy(writer, f)
	defer f.Close()
	if err2 != nil {
		log.Fatal(err2)
		return false
	}

	return true
}

func (fs *FileStorage) GetMetadata(id int) *StorageItemMetadata {
	metadata, found := fs.data[id]
	if !found {
		return nil
	}
	return metadata
}

func (fs *FileStorage) GetAllIds() []int {
	keys := make([]int, 0, len(fs.data))
	for k := range fs.data {
		keys = append(keys, k)
	}
	return keys
}

func (fs *FileStorage) HasId(id int) bool {
	return contains(fs.GetAllIds(), id)
}
