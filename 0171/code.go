//12:13
func titleToNumber(s string) int {
	var result int
	for _, char := range s {
		digit := int(char - 'A' + 1)
		result = result*26 + digit
	}
	return result
}
