package helpers

func DeleteElement[T any](slice []T, index int) ([]T, bool) {
	if index < 0 || index >= len(slice) {
		return []T{}, false
	}
	if index == 0 {
		return slice[1:], true
	}
	if index == len(slice)-1 {
		return slice[:len(slice)-1], true
	}

	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1], true
}
