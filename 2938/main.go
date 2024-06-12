func minimumSteps(s string) int64 {
	ones := 0
	ans := 0
	for _, ch := range s {
		if ch == '1' {
			ones += 1
		} else {
			ans += ones
		}
	}

	return int64(ans)
}