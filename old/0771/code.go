func numJewelsInStones(jewels string, stones string) int {
	mapJ := make(map[rune]bool)
	ans := 0

	for _, char := range jewels {
		mapJ[char] = true
	}

	for _, char := range stones {
		if mapJ[char] {
			ans++
		}
	}
	return ans
}

//13.12