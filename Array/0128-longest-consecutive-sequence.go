package main

func longestConsecutive(nums []int) int {
	numSet := make(map[int]struct{})
	for _, n := range nums {
		numSet[n] = struct{}{}
	}

	longest := 0

	for num := range numSet {
		if _, found := numSet[num-1]; !found {
			length := 1

			for {
				if _, exist := numSet[num+length]; exist {
					length++
				} else {
					break
				}
			}

			if length > longest {
				longest = length
			}
		}
	}

	return longest
}
