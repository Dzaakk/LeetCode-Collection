package _255
//42:16
func maxScoreWords(words []string, letters []byte, score []int) int {
	lettersCount:= make(map[byte]int)
	for _, letter:= range letters {
		lettersCount[letter]++
	}

	scoreMap := make(map[byte]int)
	for i, s := range score {
		scoreMap[byte('a'+i)] = s
	}

	maxScore := 0
	backtrack(words, lettersCount, scoreMap, 0, 0, &maxScore)
	return maxScore

}
func backtrack(words []string, lettersCount map[byte]int, score map[byte]int, index, currentScore int, maxScore *int)  {

	if index == len(words) || len(lettersCount) == 0 {
		if currentScore > *maxScore{
			*maxScore = currentScore
		}
		return
	}

	wordScore := 0
	wordLetterCounter := make(map[byte]int)
	isImpossibleWord := false

	for _, v := range words[index] {
		letter := byte(v)
		wordScore += score[letter]
		wordLetterCounter[letter]++
		if wordLetterCounter[letter] > lettersCount[letter] {
			isImpossibleWord = true
		}
	}
	if !isImpossibleWord {
		for letter, count := range wordLetterCounter {
			lettersCount[letter] -= count
			if lettersCount[letter] == 0 {
				delete(lettersCount, letter)
			}
		}

		// Backtrack with the next word
		backtrack(words, lettersCount, score, index+1, currentScore+wordScore, maxScore)

		// Restore available letters count
		for letter, count := range wordLetterCounter {
			lettersCount[letter] += count
		}
	}

	// Skip using the current word and move to the next
	backtrack(words, lettersCount, score, index+1, currentScore, maxScore)
}