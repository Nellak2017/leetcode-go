package internal

// TODO: Make tests for this
func IndexOf[T comparable](items []T, target T, start int) int {
	for i, v := range items[start:] {
		if v == target {
			return i + start
		}
	}
	return -1
}
