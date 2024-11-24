package kmp

type kmp []int

type Matcher interface {
	MatchStr(string, string) int
}

func BuildNextTable(needle string) Matcher {

	next := make(kmp, len(needle))
	for i, l := 1, 0; i < len(needle); {
		if needle[i] == needle[l] {
			l++
			next[i] = l
			i++
		} else {
			if l > 0 {
				l = next[l-1]
			} else {
				next[i] = 0
				i++
			}
		}
	}
	return next
}

func (k kmp) MatchStr(haystack, needle string) int {

	next := k
	i, j := 0, 0
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else if j > 0 {
			j = next[j-1]
		} else {
			i++
		}
	}

	if j == len(needle) {
		return i - j
	}

	return -1
}
