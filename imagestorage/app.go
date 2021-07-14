package imagestorage

import (
	"fmt"
	"log"
	"net/http"

	"github.com/osanj/go-image-storage/imagestorage/storage"
)

func BuildAndServe(port int) {
	// read config
	// create service
	// create controllers
	// assign controllers

	storageBackend := storage.MemoryStorage{}
	service := ImageStorageService{storage: &storageBackend}

	controllerPostImage := PostImageController{service: &service}
	controllerGetImage := GetImageController{service: &service}
	controllerListImages := GetImageListController{service: &service}

	http.HandleFunc("/upload", controllerPostImage.Serve)
	http.HandleFunc("/item", controllerGetImage.Serve)
	http.HandleFunc("/list", controllerListImages.Serve)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), nil))
}
