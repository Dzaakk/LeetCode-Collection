package main

//09:33
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[int32]int)
	for _, x := range s {
		m[x]++
	}
	for _, y := range t {
		m[y]--
	}

	for _, c := range m {
		if c > 0 {
			return false
		}
	}

	return true
}
