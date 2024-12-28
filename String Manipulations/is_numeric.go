package sprint

func IsNumeric(s string) bool {

	if len(s) == 0 {
		return false
	}

	for _, e := range s {
		if e < '0' || e > '9' {
		return false
		} 
	}
		return true

}