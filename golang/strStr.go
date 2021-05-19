func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		 return 0 
	}
	if len(needle) > len(haystack) {
			return -1
	}
	for index := 0; index < len(haystack); index++ {
			if (len(haystack) - index) < len(needle) {
					return -1
			}
			needleIndex := 0
			haystackIndex := index
			for haystack[haystackIndex] == needle[needleIndex] {
					if needleIndex == (len(needle) - 1) {
							return index // start
					}
					haystackIndex++
					needleIndex++
			}
	}
	return -1
}