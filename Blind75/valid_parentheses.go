package main

// 20:30
func isValid(s string) bool {
	m := map[byte]byte{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		c := s[i]

		if close, ok := m[c]; ok {
			stack = append(stack, close)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]

			if c != top {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
