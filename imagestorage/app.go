package imagestorage

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/osanj/go-image-storage/imagestorage/storage"
)

func BuildAndServe(configPath string, port int) {
	configuration := readConfiguration(configPath)

	var storageBackend storage.Storage
	log.Printf("serving at localhost:%d", port)
	log.Printf("using storage at %s", configuration.Storage.BasePath)
	if configuration.Storage.BasePath == BasePathMemory {
		storageBackend = storage.NewMemoryStorage()
	} else {
		storageBackend = storage.NewFileStorage(configuration.Storage.BasePath, configuration.Storage.CreateIfNotExists)
	}

	service := ImageStorageService{storage: storageBackend}

	controllerPostImage := PostImageController{service: &service}
	controllerGetImage := GetImageController{service: &service}
	controllerListImages := GetImageListController{service: &service}

	handler := RegexpHandler{}
	handler.HandleFunc(regexp.MustCompile("\\/api\\/v1\\/upload"), controllerPostImage.Serve)
	handler.HandleFunc(regexp.MustCompile("\\/api\\/v1\\/item\\/[0-9]+"), controllerGetImage.Serve)
	handler.HandleFunc(regexp.MustCompile("\\/api\\/v1\\/list"), controllerListImages.Serve)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), &handler))
}
