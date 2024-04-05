//08:59
func makeGood(s string) string {
	result := s
	good := false
	idx := 0
	for good == false {
		good = true
		ls := strings.ToLower(result)
		for i := idx; i < len(result)-1; i++ {
			if ls[i] == ls[i+1] && result[i] != result[i+1]{
				result = result[:i] + result[i+2:]
				if i > 0 {
					idx = i-1
				}
				good = false
				break
			}
		}
	}
	return result
}
