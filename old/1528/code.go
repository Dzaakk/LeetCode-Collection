//04:51
func restoreString(s string, indices []int) string {
	result := make([]rune, len(s))

	for idx , char := range s {
		result[indices[idx]] = char
	}
	return string(result)
}