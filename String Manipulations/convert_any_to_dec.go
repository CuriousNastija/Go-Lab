package sprint

import "strings"


func ConvertAnyToAny(nbr, baseFrom, baseTo string) string {

	runesNbr := []rune(nbr)
	runesBaseFrom := []rune(baseFrom)
	runesBaseTo := []rune(baseTo)
	baseFromLen := len(baseFrom)
	baseToLen := len(baseTo)
	nbrLen := len(nbr)

	if baseFromLen < 2 ||  baseToLen < 2 {
		return ""
	}
	

	for i := 0; i < baseFromLen; i++ {

		if runesBaseFrom[i] == '+' || runesBaseFrom[i] == '-' {
			return ""
		}

		for j := i + 1; j < baseFromLen; j++ {
			if runesBaseFrom[i] == runesBaseFrom[j] {
				return ""
			}
		}
	}

	for i := 0; i < baseToLen; i++ {

		if runesBaseTo[i] == '+' || runesBaseTo[i] == '-' {
			return ""
		}

		for j := i + 1; j < baseToLen; j++ {
			if runesBaseTo[i] == runesBaseTo[j] {
				return ""
			}
		}
	}

	if nbr == "" {
		return ""
	}

	var result int
	power := 1

	for i := nbrLen - 1; i >= 0; i-- {
		found := false
		for j := 0; j < baseFromLen; j++ {
			if runesNbr[i] == runesBaseFrom[j] {
				result += j * power
				power *= baseFromLen
				found = true
				break
			}
		}
		if !found {
			return ""
		}
		
	}

	if result == 0 {
		return string(runesBaseTo[0])
	}

	var converted []string
	for result > 0 {
		remainder := result % baseToLen
		converted = append(converted, string(runesBaseTo[remainder]))
		result /= baseToLen
	}

	for i, j := 0, len(converted)-1; i < j; i, j = i+1, j-1 {
		converted[i], converted[j] = converted[j], converted[i]
	}

	return strings.Join(converted, "")

}