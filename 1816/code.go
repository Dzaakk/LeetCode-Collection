//04:30
func truncateSentence(s string, k int) string {
	counter := 0
	var res string
	for _, str := range s {
		if str == ' ' {
			counter++
		}
		if counter == k {
			break
		}
		res += string(str)
	}
	return res
}

func truncateSentence(s string, k int) string {
	var last int
	for i, v := range s {
		if v == ' '{
			k--
			last = i
		}
		if k == 0{
			return s[:last]
		}
	}

	return s
}