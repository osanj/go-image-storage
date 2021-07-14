package imagestorage

type ImageElement struct {
	Id       int
	Name     string
	MimeType string
}

type ResponseImageList struct {
	Images []ImageElement
}

type ResponseImageUpload struct {
	Id int
}
