package list

import "cmp"

// Dedupe dedpulicates and returns a new list
func Dedupe[T comparable](inp []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0)
	for _, v := range inp {
		if _, ok := seen[v]; !ok {
			result = append(result, v)
			seen[v] = true
		}
	}
	return result
}

// Find intersection in two lists
func Intersection[T comparable](l1 []T, l2 []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0)
	for _, v := range l1 {
		seen[v] = true
	}

	for _, v := range l2 {
		if _, ok := seen[v]; ok {
			result = append(result, v)
		}
	}

	return result
}

func Frequency[T comparable](l []T) map[T]int {
	freq := map[T]int{}
	for _, item := range l {
		freq[item] = freq[item] + 1
	}
	return freq
}

func Sum[T cmp.Ordered](l []T) T {
	var r T

	for _, v := range l {
		r += v
	}

	return r
}
