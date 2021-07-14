package storage

type Storage interface {
	put(data []byte, name string) int
	get(id int) []byte
	getAllIds() []int
	hasId(id int) bool
}
