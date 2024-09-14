package sliceutil

// Filter filters a slice based on a given function.
// It returns a new slice containing only the elements that satisfy the function.
func Filter[T any](slice []T, f func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map maps a slice based on a given function.
// It returns a new slice containing the results of applying the function to each element.
func Map[T, R any](slice []T, f func(T) R) []R {
	var result []R
	for _, v := range slice {
		result = append(result, f(v))
	}
	return result
}

// GroupBy groups a slice based on a given function.
// It returns a map where the keys are the results of applying the function to each element,
func GroupBy[T any, R comparable](slice []T, f func(T) R) map[R][]T {
	result := make(map[R][]T)
	for _, v := range slice {
		key := f(v)
		result[key] = append(result[key], v)
	}
	return result
}

// Contains returns true if the given value is present in the slice.
// It uses the == operator to compare the values.
func Contains[T comparable](slice []T, v T) bool {
	for _, x := range slice {
		if x == v {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of the given value in the slice.
// It uses the == operator to compare the values.
// If the value is not present in the slice, it returns -1.
func IndexOf[T comparable](slice []T, v T) int {
	for i, x := range slice {
		if x == v {
			return i
		}
	}
	return -1
}

// Remove removes all occurrences of the given value from the slice.
// It uses the == operator to compare the values.
// It returns a new slice with the removed elements.
func Remove[T comparable](slice []T, v T) []T {
	var result []T
	for _, x := range slice {
		if x != v {
			result = append(result, x)
		}
	}
	return result
}

// RemoveAll is an alias for Remove.
// It is provided for consistency with the standard library.
func RemoveAll[T comparable](slice []T, v T) []T {
	return Remove(slice, v)
}

// Reverse reverses the order of the elements in the slice.
// It modifies the slice in place.
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
