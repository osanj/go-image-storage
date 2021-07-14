package imagestorage

import (
	"fmt"
	"mime"
	"net/http"
	"strings"
)

func writeHttpError(w http.ResponseWriter, status int, msg *string) {
	w.WriteHeader(status)
	if msg == nil {
		fmt.Fprintf(w, "%d", status)
	} else {
		fmt.Fprintf(w, "%d - %s!", status, *msg)
	}
}

// https://gist.github.com/rjz/fe283b02cbaa50c5991e1ba921adf7c9
func hasContentType(r *http.Request, mimetype string) bool {
	contentType := r.Header.Get("content-type")
	if contentType == "" {
		return mimetype == "application/octet-stream"
	}

	for _, v := range strings.Split(contentType, ",") {
		t, _, err := mime.ParseMediaType(v)
		if err != nil {
			break
		}
		if t == mimetype {
			return true
		}
	}
	return false
}

type PostImageController struct {
	service *ImageStorageService
}

func (h *PostImageController) Serve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		msg := "not a " + http.MethodPost + " request"
		writeHttpError(w, http.StatusMethodNotAllowed, &msg)
		return
	}

	allowedMimeTypes := []string{"image/jpg", "image/jpeg", "image/png"}
	mimeTypeNone := ""
	mimeType := mimeTypeNone

	for _, allowedMimeType := range allowedMimeTypes {
		if hasContentType(r, allowedMimeType) {
			mimeType = allowedMimeType
			break
		}
	}

	if mimeType == mimeTypeNone {
		msg := "allowed mime types are: " + strings.Join(allowedMimeTypes, ", ")
		writeHttpError(w, http.StatusUnsupportedMediaType, &msg)
		return
	}

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

type GetImageController struct {
	service *ImageStorageService
}

func (h *GetImageController) Serve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		msg := "not a " + http.MethodGet + " request"
		writeHttpError(w, http.StatusMethodNotAllowed, &msg)
		return
	}

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

}

type GetImageListController struct {
	service *ImageStorageService
}

func (h *GetImageListController) Serve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		msg := "not a " + http.MethodGet + " request"
		writeHttpError(w, http.StatusMethodNotAllowed, &msg)
		return
	}

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
