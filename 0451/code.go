//24:13
type CharFrequency struct {
	Char byte
	Freq int
}
func frequencySort(s string) string {
	freqMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freqMap[s[i]]++
	}
	charFreqs := make([]CharFrequency, 0 , len(freqMap))
	for char, freq := range freqMap {
		charFreqs = append(charFreqs, CharFrequency{char, freq})
	}

	sort.Slice(charFreqs, func(i, j int) bool {
		if charFreqs[i].Freq == charFreqs[j].Freq {
			return charFreqs[i].Char < charFreqs[j].Char
		}
		return charFreqs[i].Freq > charFreqs[j].Freq
	})

	result := make([]byte, 0, len(s))
	for _, cf := range charFreqs{
		for i := 0; i < cf.Freq; i++ {
			result = append(result, cf.Char)
		}
	}

	return string(result)
}
