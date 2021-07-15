package imagestorage

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/osanj/go-image-storage/imagestorage/storage"
)

func BuildAndServe(port int) {
	// read config
	// create service
	// create controllers
	// assign controllers

	// storageBackend := storage.NewMemoryStorage()
	storageBackend := storage.NewFileStorage("/home/jonas/code/personal/go-image-storage/test")
	service := ImageStorageService{storage: storageBackend}

	controllerPostImage := PostImageController{service: &service}
	controllerGetImage := GetImageController{service: &service}
	controllerListImages := GetImageListController{service: &service}

	handler := RegexpHandler{}
	handler.HandleFunc(regexp.MustCompile("\\/upload"), controllerPostImage.Serve)
	handler.HandleFunc(regexp.MustCompile("\\/item\\/[0-9]+"), controllerGetImage.Serve)
	handler.HandleFunc(regexp.MustCompile("\\/list"), controllerListImages.Serve)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), &handler))
}
