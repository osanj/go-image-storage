package imagestorage

// import (
//     "encoding/json"
//     "os"
//     "fmt"
// )

// type Configuration struct {
//     Users    []string
//     Groups   []string
// }

// file, _ := os.Open("conf.json")
// defer file.Close()
// decoder := json.NewDecoder(file)
// configuration := Configuration{}
// err := decoder.Decode(&configuration)
// if err != nil {
//   fmt.Println("error:", err)
// }
// fmt.Println(configuration.Users) // output: [UserA, UserB]

// https://stackoverflow.com/questions/16465705/how-to-handle-configuration-in-go
