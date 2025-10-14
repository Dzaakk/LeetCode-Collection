package main

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		if !isAlphaOrNum(s[l]) {
			l++
			continue
		}
		if !isAlphaOrNum(s[r]) {
			r--
			continue
		}

		if toLowerASCII(s[l]) != toLowerASCII(s[r]) {
			return false
		}
		l++
		r--
	}

	return true
}

func isAlphaOrNum(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

func toLowerASCII(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
