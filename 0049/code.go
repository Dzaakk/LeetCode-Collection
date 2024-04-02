//11:31
func groupAnagrams(strs []string) [][]string {
	keyAnagrams := make(map[string][]string)

	for _, str := range strs {
		chars := []byte(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		key := string(chars)
		keyAnagrams[key] = append(keyAnagrams[key], str)
	}

	var result [][]string
	for _, anagram := range keyAnagrams {
		result = append(result, anagram)
	}

	return result
}

