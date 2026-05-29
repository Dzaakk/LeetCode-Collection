package main

//18:45
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	result := []int{}
	counter := 0

	for _, n := range nums {
		m[n]++

		if counter < m[n] {
			counter = m[n]
		}
	}

	for k > 0 {
		for key, value := range m {
			if value >= counter {
				result = append(result, key)
				m[key] = 0
				k--
			}
		}
		counter--
	}

	return result
}
