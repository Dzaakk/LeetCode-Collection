//44:33
func minimumPushes(word string) int {
	freq := make([]int, 26)
	for _, char := range word {
		freq[char - 'a']++
	}

	slices.Sort(freq)
	result := 0
	multiplier := 1
	counter := 8
	for i := 25; i >= 0; i--{
		res += freq[i] * multiplier
		counter--
		if counter == 0 {
			multiplier++
			counter = 8
		}
	}
	return result
}
func minimumPushes(word string) int {
	counter := make([]int, 26)

	for _, char := range word {
		counter[char - 'a']++
	}

	slices.SortFunc(counter, func(a, b int) int {
		return b-a
	})
	result := 0
	for i := 0; i < len(counter); i++ {
		multi := i / 8 + 1
		result += counter[i] * multi
	}
	return result
}
