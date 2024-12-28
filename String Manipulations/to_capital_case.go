package sprint

func ToCapitalCase(s string) string {

	strCapitalCase := ""
	isStartOfWord := true

	for _, e := range s {

		if (e >= 'A' && e <= 'Z') || (e >= 'a' && e <= 'z') {
			if isStartOfWord {
				if e >= 'a' && e <= 'z' {
					e = e - 32
				}

				strCapitalCase += string(e)

			} else {
				if e >= 'A' && e <= 'Z' {
					e = e + 32
				}
				strCapitalCase += string(e)
			}
			isStartOfWord = false
		} else if e >= '0' && e <= '9' {
			strCapitalCase += string(e)
			isStartOfWord = false
		} else {
			strCapitalCase += string(e)
			isStartOfWord = true
		}
	}

	return strCapitalCase

}