package main

//08:23
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	result := [][]string{}
	for _, str := range strs {
		key := generateKey(str)

		m[key] = append(m[key], str)
	}

	for _, v := range m {
		result = append(result, v)
	}

	return result
}

func generateKey(s string) [26]int {
	keys := [26]int{}
	for _, c := range s {
		keys[c-'a']++
	}

	return keys
}
