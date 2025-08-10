func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	summary := []string{}
	start := 0

	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 || nums[i]+1 != nums[i+1] {
			if start == i {
				summary = append(summary, fmt.Sprintf("%d", nums[start]))
			} else {
				summary = append(summary, fmt.Sprintf("%d->%d", nums[start], nums[i]))
			}
			start = i + 1
		}
	}

	return summary
}