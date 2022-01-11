package functions

import "strconv"

func Power(base int, powerof int) int {
	result := 1

	for powerof > 0 {
		result *= base
		powerof--
	}

	return result
}

func Hex(s string) string {
	result := 0
	split := []rune(s)
	

	for i, j := range split {
		powerof := len(split) - (i + 1)
		var multiplier int

		if powerof > 0 {
			multiplier = Power(16, powerof)
		} else {
			multiplier = 1
		}

		if j > 47 && j < 58 {
			result += int(j - 48) * multiplier
		} else if j > 64 && j < 71 {
			temp := int(j - 55)

			result += temp * multiplier
		}
	}

	return strconv.Itoa(result)
}

func Bin(s string) string {
	result := 0
	split := []rune(s)
	

	for i, j := range split {
		powerof := len(split) - (i + 1)
		var multiplier int

		if powerof > 0 {
			multiplier = Power(2, powerof)
		} else {
			multiplier = 1
		}

		if j == 49 {
			result += multiplier
		}
	}

	return strconv.Itoa(result)
}