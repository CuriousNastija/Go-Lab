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

func IsSorted(f func(a, b string) int, arr []string) bool {

	n := len(arr)

	if n <= 1 {
		return true
	}

	ascending := true
	descending := true

	for i := 0; i < n-1; i++ {
		if f(arr[i], arr[i+1]) > 0 {
			ascending = false
		}
		if f(arr[i], arr[i+1]) < 0 {
			descending = false
		}
	}

	return ascending || descending

}