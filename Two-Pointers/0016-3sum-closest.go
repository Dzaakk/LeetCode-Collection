package main

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := n - 1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if abs(sum-target) < abs(closest-target) {
				closest = sum
			}

			if sum == target {
				return sum
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
	}

	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
