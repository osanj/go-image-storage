package main

import (
	"os"
	"strconv"

	"github.com/osanj/go-image-storage/imagestorage"
)

func main() {
	if len(os.Args) != 3 {
		panic("please provide a config path and port")
	}

	configPath := os.Args[1]
	port, err := strconv.Atoi(os.Args[2])

	if err != nil {
		panic("please provide a valid port")
	}

	imagestorage.BuildAppAndServe(configPath, port)
}
