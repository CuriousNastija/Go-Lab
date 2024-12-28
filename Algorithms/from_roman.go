package sprint

func GetValue(ch rune) int {

	switch ch {
		case 'I':
			return 1
		case 'V':
			return 5
		case 'X':
			return 10
		case 'L':
			return 50
		case 'C':
			return 100
		case 'D':
			return 500
		case 'M':
			return 1000
		default:
			return 0
		}
}

func FromRoman(s string) int {

	var result int

	for i := 0; i < len(s); i++ {
		current := GetValue(rune(s[i]))

		if i < len(s)-1 && current < GetValue(rune(s[i+1])) {
			result -= current
		} else {
			result += current
		}
	}
	
	return result

}