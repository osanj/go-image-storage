package storage

// https://stackoverflow.com/a/10485970
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
