package list

import "cmp"

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

func Sum[T cmp.Ordered](l []T) T {
	result := T(0)
	for _, v := range l {
		result += v
	}

	return result
}

func Min[T cmp.Ordered](l []T) (T, int) {
	if len(l) < 1 {
		panic("list must not be empty")
	}
	idx := 0
	result := l[0]
	for i, v := range l {
		if v < result {
			result = v
			idx = i
		}
	}
	return result, idx
}

func Frequency[T comparable](l []T) map[T]int {
	freq := map[T]int{}
	for _, item := range l {
		if count, found := freq[item]; found {
			freq[item] = count + 1
		} else {
			freq[item] = 1
		}
	}
	return freq
}
