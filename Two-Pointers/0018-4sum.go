package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}

	n := len(nums)
	if n < 4 {
		return res
	}

	sort.Ints(nums)

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}

		if nums[i]+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}

			if nums[i]+nums[j]+nums[n-1]+nums[n-2] < target {
				continue
			}

			l := j + 1
			r := n - 1

			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]

				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--

					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				} else if sum < target {
					l++
				} else {
					r--
				}
			}

		}
	}
	return res
}
