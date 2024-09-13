package sliceutil

func Filter[T any](slice []T, f func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map[T, R any](slice []T, f func(T) R) []R {
	var result []R
	for _, v := range slice {
		result = append(result, f(v))
	}
	return result
}
