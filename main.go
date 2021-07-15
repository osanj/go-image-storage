package main

import (
	"os"

	"github.com/osanj/go-image-storage/imagestorage"
)

func main() {
	if len(os.Args) != 2 {
		panic("Please provide a config path")
	}
	imagestorage.BuildAndServe(os.Args[1], 8080)
}
