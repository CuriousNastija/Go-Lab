package sprint

func IsLower(s string) bool {

	if len(s) == 0 {
		return false
	}

	for _, e := range s {
		if e < 'a' || e > 'z' {
		return false
		} 
	}
		return true
}