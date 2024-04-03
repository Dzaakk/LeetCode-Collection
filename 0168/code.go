//18:40
func convertToTitle(columnNumber int) string {
	if columnNumber <= 26 {
		return string('A'+(columnNumber-1))
	} else if columnNumber%26 == 0 {
		return convertToTitle((columnNumber-1)/26) + "Z"
	}
	return convertToTitle(columnNumber/26) + convertToTitle(columnNumber%26)
}
