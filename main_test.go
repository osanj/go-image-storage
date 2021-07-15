package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/osanj/go-image-storage/imagestorage"
)

func GetListOfImages(t *testing.T, host string) *imagestorage.ResponseImageList {
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/list", host))
	if err != nil {
		t.Fatal("could not get list of items")
	}
	defer resp.Body.Close()
	var respData imagestorage.ResponseImageList
	json.NewDecoder(resp.Body).Decode(&respData)
	return &respData
}

func UploadImage(t *testing.T, host string, mimeType string, reader io.Reader) *imagestorage.ResponseImageUpload {
	resp, err := http.Post(fmt.Sprintf("%s/api/v1/upload", host), mimeType, reader)
	if err != nil {
		t.Fatal("could upload image")
	}
	defer resp.Body.Close()
	var respData imagestorage.ResponseImageUpload
	json.NewDecoder(resp.Body).Decode(&respData)
	return &respData
}

func DownloadImage(t *testing.T, host string, id int) io.Reader {
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/item/%d", host, id))
	if err != nil {
		t.Fatal("could upload image")
	}
	defer resp.Body.Close()
	return resp.Body
}

func LoadFile(path string) io.Reader {
	f, err := os.Open(path)
	if err != nil {
		panic("Could not open file")
	}
	return f
}

// https://gist.github.com/dixudx/3989284b142414e10352fde9def5c771
func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

func TestHappyPath(t *testing.T) {
	handler := imagestorage.BuildApp("config/dev.json")
	server := httptest.NewServer(handler)
	defer server.Close()

	// check that service is empty
	imageList1 := GetListOfImages(t, server.URL)
	if len(imageList1.Images) != 0 {
		t.Fatal("store should be empty at the beginning")
	}

	// add images
	f1 := LoadFile("testdata/beaver1.jpg")
	f2 := LoadFile("testdata/beaver2.jpg")

	upload1 := UploadImage(t, server.URL, "image/jpeg", f1)
	if upload1.Id != 1 {
		t.Fatal("unexpected id for first image")
	}
	upload2 := UploadImage(t, server.URL, "image/jpeg", f2)
	if upload2.Id != 2 {
		t.Fatal("unexpected id for second image")
	}

	// check that service is populated
	imageList2 := GetListOfImages(t, server.URL)
	if len(imageList2.Images) != 2 {
		t.Fatal("store should have two items")
	}

	// check image data
	expBytes1 := StreamToByte(f1)
	expBytes2 := StreamToByte(f2)
	actBytes1 := StreamToByte(DownloadImage(t, server.URL, 1))
	actBytes2 := StreamToByte(DownloadImage(t, server.URL, 2))
	if !bytes.Equal(expBytes1, actBytes1) {
		t.Fatal("stored image 1 is differing from original")
	}
	if !bytes.Equal(expBytes2, actBytes2) {
		t.Fatal("stored image 2 is differing from original")
	}
}
