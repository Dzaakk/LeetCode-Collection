package main

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	l, r := 0, 0
	counter := 0

	for r < len(t) && l < len(s) {
		if s[l] == t[r] {
			l++
			counter++
		}
		r++
	}
	return counter == len(s)
}
