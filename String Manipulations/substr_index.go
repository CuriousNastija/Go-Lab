package sprint

func SubstrIndex(s string, toFind string) int {

	slen := len(s)
	toFindLen := len(toFind)

	if toFindLen == 0 {
		return 0
	}


	for i := 0; i <= slen - toFindLen; i++ {
		if s[i:i+toFindLen] == toFind {
			return i
		} 
	}

	return -1

}