package imagestorage

import (
	"fmt"
	"net/http"
)

type PostImageController struct {
	service *ImageStorageService
}

func (h *PostImageController) Serve(w http.ResponseWriter, r *http.Request) {
	// w.Write(h.Thing)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

type GetImageController struct {
	service *ImageStorageService
}

func (h *GetImageController) Serve(w http.ResponseWriter, r *http.Request) {
	// w.Write(h.Thing)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

type GetImageListController struct {
	service *ImageStorageService
}

func (h *GetImageListController) Serve(w http.ResponseWriter, r *http.Request) {
	// w.Write(h.Thing)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
