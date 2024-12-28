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

func IsUpper(s string) bool {

	if len(s) == 0 {
		return false
	}

	for _, e := range s {
		if e < 'A' || e > 'Z' {
		return false
		} 
	}
		return true


}

func IsAlphanumeric(s string) bool {

	if len(s) == 0 {
        return false
    }

    for _, e := range s {
        if !(e >= '0' && e <= '9') && !(e >= 'A' && e <= 'Z') && !(e >= 'a' && e <= 'z') {
            return false
        }
    }
    return true

}

func ArrCountIf(f func(string) bool, a []string) int {

	count := 0
    for _, str := range a {
        if f(str) {
            count++
        }
    }
    return count
}