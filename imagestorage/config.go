package imagestorage

import (
	"encoding/json"
	"os"
)

const BasePathMemory = ":memory:"

type StorageConfiguration struct {
	BasePath          string
	CreateIfNotExists bool
}

type Configuration struct {
	Storage StorageConfiguration
}

func readConfiguration(path string) Configuration {
	file, _ := os.Open(path)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}
	return configuration
}
