//18:30
func intToRoman(num int) string {
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := []string{}

	for i := 0; i < len(symbols); i++ {
		for num >= values[i] {
			num -= values[i]
			result = append(result, symbols[i])
		}
	}

	return strings.Join(result, "")
}
