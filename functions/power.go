package functions

func Power(base int, powerof int) int {
	result := 1

	for powerof > 0 {
		result *= base
		powerof--
	}

	return result
}