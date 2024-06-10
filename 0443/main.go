func compress(chars []byte) int {
	var count, index int = 1, 0
	for i := 1; i <= len(chars); i++ {
		if i < len(chars) && chars[i] == chars[i-1] {
			count++
		} else {
			chars[index] = chars[i-1]
			index++
			if count > 1 {
				for _, ch := range strconv.Itoa(count) {
					chars[index] = byte(ch)
					index++
				}
			}
			count = 1
		}
	}
	return index
}