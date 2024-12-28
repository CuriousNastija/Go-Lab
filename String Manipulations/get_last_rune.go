package sprint

func GetLastRune(s string) rune {

	runeArray := []rune(s)

	if len(runeArray) > 0 {
		return runeArray[len(runeArray) - 1]
	}
	
	return 0

}