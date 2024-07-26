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

func intToRoman(num int) string {
	var res string
	roman := map[int]string{
		1: "I",
		5:"V",
		10:"X",
		50:"L",
		100:"C",
		500:"D",
		1000:"M",
	}

	for num != 0 {
		switch {
		case num - 1000 >= 0:
			res += roman[1000]
			num -= 1000
		case num-900 >= 0:
			res += roman[100] + roman[1000]
			num -= 900
		case num - 500 >= 0:
			res += roman[500]
			num -= 500
		case num - 400 >= 0:
			res += roman[100] + roman[500]
			num -= 400
		case num - 100 >= 0:
			res += roman[100]
			num -= 100
		case num - 90 >= 0:
			res += roman[10] + roman[100]
			num -= 90
		case num - 50 >= 0:
			res += roman[50]
			num -= 50
		case num - 40 >= 0:
			res += roman[10] + roman[50]
			num -= 40
		case num - 10 >= 0:
			res += roman[10]
			num -= 10
		case num - 9 >= 0:
			res += roman[1] + roman[10]
			num -= 9
		case num - 5 >= 0:
			res += roman[5]
			num -= 5
		case num - 4 >= 0:
			res += roman[1] + roman[5]
			num -= 4
		case num - 1 >= 0:
			res += roman[1]
			num -= 1
		}
	}
	return res
}