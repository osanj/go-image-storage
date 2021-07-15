# go-image-storage

`go run main.go config/prod.json`

`go test -v`

`curl -X POST -H "Content-Type: image/jpeg" --data-binary @testdata/beaver1.jpg localhost:8080/upload`

`curl localhost:8080/item/0 -o /dev/null -i`



https://gobyexample.com
https://golang.org/doc/
https://golang.org/doc/articles/wiki/
https://github.com/ybkuroki/go-webapp-sample
