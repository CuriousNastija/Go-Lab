package sprint

func StrCompare(a, b string) int {

	runesA := []rune(a)
	runesB := []rune(b)
	
	var compare int

	if len(runesA) == 0 && len(runesB) == 0 {
		return 0
	}
	if len(runesA) == 0 && len(runesB) != 0 {
		return -1
	}

	if len(runesA) != 0 && len(runesB) == 0 {
		return 1
	}

	for i := 0; i < len(runesA); i++ {
		compare = int(runesA[i] - runesB[i])
		if compare < 0 {
			return -1
		} else if compare > 0 {
			return 1
		} else {
			if len(runesA) < len(runesB) {
				return -1
			} else if len(runesA) > len(runesB) {
				return 1
			} else {
				return 0
			}
		}
	} 
	
	return compare
}