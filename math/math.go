package math

// GCD is greatest common divisior
func GCD(a, b int) int {
	result := 0
	if a > b {
		result = b
	} else {
		result = a
	}

	for result > 0 {
		if a%result == 0 && b%result == 0 {
			return result
		}
		result--
	}

	return result
}

// LCM is least common multiple
func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}
