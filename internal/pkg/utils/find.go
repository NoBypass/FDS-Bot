package utils

func Find[T any](slice []*T, predicate func(*T) bool) *T {
	for _, v := range slice {
		if predicate(v) {
			return v
		}
	}
	return nil
}
