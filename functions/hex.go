package functions

import "strconv"

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