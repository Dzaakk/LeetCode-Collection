func strStr(haystack string, needle string) int {
  check := false 
		
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			for j := 0; j < len(needle); j++ {
				if i +j < len(haystack) && haystack[i+j] == needle[j] {
					check = true  
				} else {
					check = false 
					break;
				}
			}
		}
		if check == true{
			return i
		}
	}
	return -1
}