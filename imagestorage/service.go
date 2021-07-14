package imagestorage

import (
	"io"
	"log"

	"github.com/osanj/go-image-storage/imagestorage/storage"
)

type ImageStorageService struct {
	storage storage.Storage
}

func (s *ImageStorageService) GetListOfImages() []ImageElement {
	elements := []ImageElement{}
	for _, id := range s.storage.GetAllIds() {
		metadata := s.storage.GetMetadata(id)
		elements = append(elements, ImageElement{Id: id, Name: metadata.Name, MimeType: metadata.MimeType})
	}
	return elements
}

func (s *ImageStorageService) HasImage(id int) bool {
	return s.storage.HasId(id)
}

func (s *ImageStorageService) GetImage(id int, writer io.Writer) bool {
	return s.storage.GetBytes(id, writer)
}

func (s *ImageStorageService) GetImageMetadata(id int) *storage.StorageItemMetadata {
	return s.storage.GetMetadata(id)
}

func (s *ImageStorageService) PutImage(reader io.Reader, name string, mimeType string) int {
	metadata := storage.StorageItemMetadata{Name: name, MimeType: mimeType}
	id := s.storage.Put(reader, metadata)
	log.Printf("image was pushed id=%d name=%s", id, name)
	return id
}
