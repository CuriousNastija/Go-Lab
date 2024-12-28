package sprint

func StrReverse(s string) string {

	runes := []rune(s)
	reversedRunes := make([]rune, len(runes))

	for i, r := range runes {
		reversedRunes[len(runes)-1-i] = r
	}
	
	return string(reversedRunes)

}