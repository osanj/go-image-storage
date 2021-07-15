package imagestorage

import (
	"io"
	"log"
	"strings"

	"github.com/osanj/go-image-storage/imagestorage/storage"
)

type ImageStorageService struct {
	storage storage.Storage
}

func (s *ImageStorageService) GetSupportedMimeTypes() []string {
	return []string{"image/jpeg", "image/png"}
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

func (s *ImageStorageService) PutImage(reader io.Reader, mimeType string) int {
	ext := strings.Split(mimeType, "/")[1]
	name := "image." + ext
	metadata := storage.StorageItemMetadata{Name: name, MimeType: mimeType}
	id := s.storage.Put(reader, metadata)
	log.Printf("an image was pushed (id=%d)", id)
	return id
}
