func smallerNumbersThanCurrent(nums []int) []int {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	counts := make([]int, len(nums))
	for i, num := range nums {
		counts[i] = sort.Search(len(sortedNums), func(j int) bool {
			return num <= sortedNums[j]
		})
	}
	return
}

// 02.16

func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {

		counter := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] < nums[i] {
				counter++
			}
		}
		res[i] = counter
	}
	return res
}

// 06.08