package sprint

func GetFirstRune(s string) rune {

	runeArray := []rune(s)

	if len(runeArray) > 0 {
		return runeArray[0]
	}
	
	return 0
}