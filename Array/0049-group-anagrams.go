package main

func groupAnagrams(words []string) [][]string {
	mapGroup := make(map[[26]int][]string)

	for _, word := range words {
		var count [26]int
		for _, c := range word {
			count[c-'a']++
		}
		mapGroup[count] = append(mapGroup[count], word)
	}
	var result [][]string
	for _, group := range mapGroup {
		result = append(result, group)
	}

	return result
}
