package sprint

func SplitWhitespaces(s string) []string {

	if len(s) == 0 {
		return []string{}
	}
	
	result := []string{}
	word := ""

	for _, e := range s {
		if e != ' ' && e != '\t' && e != '\n' && e != ',' {
			word += string(e)
		} else {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		}
	}
	if word != "" {
		result = append(result, word)
	}

	return result

}