package sprint

func StrCompare(a, b string) int {

	runesA := []rune(a)
	runesB := []rune(b)

	minLen := len(runesA)
	if len(runesB) < minLen {
		minLen = len(runesB)
	}

	for i := 0; i < minLen; i++ {
		if runesA[i] < runesB[i] {
			return -1
		} else if runesA[i] > runesB[i] {
			return 1
		}
	} 

	if len(runesA) < len(runesB) {
		return -1
	} else if len(runesA) > len(runesB) {
		return 1
	}
	
	return 0
}

func AdvancedSortWordArr(a []string, f func(a, b string) int) []string {

	n := len(a)
	
	for i := 0; i < len(a); i++ {
		for j := 0; j < n-1-i; j++ {
			if f(a[j], a[j+1]) > 0 {
			a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}