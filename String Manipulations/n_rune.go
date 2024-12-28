package sprint

func NRune(s string, i int) rune {

	runeArray := []rune(s)

	if len(runeArray) > 0 {
		return runeArray[i]
	}
	
	return 0


}