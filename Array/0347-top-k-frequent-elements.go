package main

func topKFrequent(nums []int, k int) []int {
	if len(nums) == k {
		return nums
	}

	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	buckets := make([][]int, len(nums)+1)
	for n, f := range freq {
		buckets[f] = append(buckets[f], n)
	}

	res := make([]int, 0, k)
	for f := len(buckets) - 1; f >= 1 && len(res) < k; f-- {
		for _, n := range buckets[f] {
			res = append(res, n)
			if len(res) == k {
				break
			}
		}
	}
	return res
}
