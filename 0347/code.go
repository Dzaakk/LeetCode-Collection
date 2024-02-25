package _347
//18:03
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	res := make([]int, k)
	for _, num := range nums {
		m[num]++
	}

	for i := 0; i < k; i++ {
		counter := 0
		val := 0
		for k, v:= range m {
			if v > counter{
				counter = v
				val = k
			}
		}
		res[i] = val
		delete(m, val)
	}
	return res

}
