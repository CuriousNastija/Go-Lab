package sprint

/* Create a function that takes two strings
and a delimiter. The function should combine
the two strings, placing the delimiter between them,
and return the combined result as a single string. */

func StrConcat(s1, s2, delim string) string {
	res := s1 + delim + s2
	return res
}