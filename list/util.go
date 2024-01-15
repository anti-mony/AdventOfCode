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
