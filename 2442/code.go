//23:06
func countDistinctIntegers(nums []int) int {
	distinct := make(map[int]bool)

	for _, num := range nums {
		reversed := reverseDigits(num)
		distinct[num] = true
		distinct[reversed] = true
	}

	return len(distinct)
}

func reverseDigits(num int) int {
	reversed := 0

	for num > 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		num /= 10
	}

	return reversed
}

