package sprint

func ToUpperCase(s string) string {

	strUpperCase := ""

	for _, e := range s {
		if e >= 'a' && e <= 'z' {
		strUpperCase += string(e - 32) 
		} else {
			strUpperCase += string(e)
		}
	}
return strUpperCase
}