func isPalindrome(x int) bool {
	numString := strconv.Itoa(x)
	n := len(numString)
	for i := 0; i < len(numString)/2; i++ {
		if rune(numString[i]) != rune(numString[n-1-i]) {
			return false
		}
	}
	return true
}