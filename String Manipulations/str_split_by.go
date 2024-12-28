package sprint

func StrSplitBy(s, sep string) []string {

	result := []string{}
	seplen := len(sep)
	slen := len(s)
	start := 0

	if len(s) == 0 {
		return []string{}
	}
	
	for i := 0; i < slen - seplen; {
		if s[i:i+seplen] == sep {
			result = append(result, s[start:i])
			i += seplen
			start = i
		} else {
			i++
		}
	}

	if start < slen {
		result = append(result, s[start:])
	}

	return result

}