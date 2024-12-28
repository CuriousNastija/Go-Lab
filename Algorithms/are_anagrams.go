package sprint

import (
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {

	a := CleanAndSortString(str1)
	b := CleanAndSortString(str2)

	return a == b

}

func CleanAndSortString(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)

	runes := []rune(s)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}