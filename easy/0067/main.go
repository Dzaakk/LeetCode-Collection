//23:31
func addBinary(a string, b string) string {
	result := ""
	carry := 0
	i, j := len(a)-1, len(b)-1

	for i >= 0 || j >= 0 || carry > 0 {
		digitA, digitB := 0 , 0

		if i >= 0 {
			digitA = int(a[i] - '0')
			i--
		}
		if j >= 0 {
			digitB = int(b[j] - '0')
			j--
		}

		currentSum := digitA + digitB + carry
		result = fmt.Sprintf("%d%s", currentSum%2, result)
		carry = currentSum/2
	}
	return result
}