package list

// Dedupe dedpulicates and returns a new list
func Dedupe[T string | int | float64 | float32](inp []T) []T {
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
func Intersection[T string | int | float64 | rune](l1 []T, l2 []T) []T {
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
