package sprint

func NbrBase(n int, base string) string {

	runesBase := []rune(base)
	baseLen := len(runesBase)

	if baseLen < 2 {
		return "NV"
	}

	for i := 0; i < baseLen; i++ {

		if runesBase[i] == '+' || runesBase[i] == '-' {
			return "NV"
		}
		
		for j := i + 1; j < baseLen; j++ {
			if runesBase[i] == runesBase[j] {
				return "NV"
			}
		}
	}

	if n == 0 {
		return string(runesBase[0])
	}

	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}

	var digits []rune
	for n > 0 {
		digit := n % baseLen
		digits = append(digits, runesBase[digit])
		n /= baseLen
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i + 1, j - 1 {
			digits[i], digits[j] = digits[j], digits[i]
	}
	
	result := string(digits)


	if isNegative {
		result = "-" + result
	}
	return result
	
}