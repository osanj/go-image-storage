# go-image-storage

Simple image storage service. Can be configured to store images on disk or in memory (for testing).
User management and authentication is not implemented.

Run with `go run main.go <configPath> <port>`, e.g. `go run main.go config/prod.json 8080`

Test with `go test -v`


## API

### Endpoints

```
POST /api/v1/upload
GET /api/v1/item/:id
GET /api/v1/list
```

### Example

```bash
curl -X POST -H "Content-Type: image/jpeg" --data-binary @testdata/beaver1.jpg http://localhost:8080/api/v1/upload
curl http://localhost:8080/api/v1/list
curl http://localhost:8080/api/v1/item/1 -o img.jpeg && eog img.jpeg
```
