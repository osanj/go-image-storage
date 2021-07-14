package imagestorage

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"strconv"
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

func writeHttpErrorNoMsg(w http.ResponseWriter, status int) {
	writeHttpError(w, status, nil)
}

func writeHttpErrorWithMsg(w http.ResponseWriter, status int, msg string) {
	writeHttpError(w, status, &msg)
}

func writeJsonResponse(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
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

func (c *PostImageController) Serve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeHttpErrorWithMsg(w, http.StatusMethodNotAllowed, "not a "+http.MethodPost+" request")
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
		writeHttpErrorWithMsg(w, http.StatusUnsupportedMediaType, "allowed mime types are: "+strings.Join(allowedMimeTypes, ", "))
		return
	}

	if r.Body == nil {
		writeHttpErrorWithMsg(w, http.StatusBadRequest, "body must not be empty")
		return
	}

	id := c.service.PutImage(r.Body, "abc", mimeType)
	status := http.StatusOK
	if status < 0 {
		status = http.StatusInternalServerError
	}
	writeJsonResponse(w, status, ResponseImageUpload{Id: id})
}

type GetImageController struct {
	service *ImageStorageService
}

func (c *GetImageController) Serve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeHttpErrorWithMsg(w, http.StatusMethodNotAllowed, "not a "+http.MethodGet+" request")
		return
	}

	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) == 0 {
		writeHttpErrorNoMsg(w, http.StatusNotFound)
	}

	id, err := strconv.Atoi(urlParts[len(urlParts)-1])
	if err != nil {
		writeHttpErrorWithMsg(w, http.StatusBadRequest, "cannot parse id")
		return
	}
	if !c.service.HasImage(id) {
		writeHttpErrorWithMsg(w, http.StatusNotFound, "requested image not found")
		return
	}

	w.Header().Set("Content-Type", c.service.GetImageMetadata(id).MimeType)
	w.WriteHeader(http.StatusOK)
	c.service.GetImage(id, w)
}

type GetImageListController struct {
	service *ImageStorageService
}

func (c *GetImageListController) Serve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeHttpErrorWithMsg(w, http.StatusMethodNotAllowed, "not a "+http.MethodGet+" request")
		return
	}

	images := c.service.GetListOfImages()
	writeJsonResponse(w, http.StatusOK, ResponseImageList{Images: images})
}
