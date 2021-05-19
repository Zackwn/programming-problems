func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	lngCmnPrx := ""
	prefixCount := 1

MainLoop:
	for {
		lastStrPrefix := ""
		for _, str := range strs {
			// check insufficient string
			if len(str) < prefixCount {
				break MainLoop
			}
			// check if there is no last str prefix
			if len(lastStrPrefix) == 0 {
				lastStrPrefix = str[0:prefixCount]
				continue
			}
			// see if the prefixes are equal
			if (str[0:prefixCount]) != lastStrPrefix {
				break MainLoop
			}
			// prefixes are equal
			lastStrPrefix = str[0:prefixCount]
		}
		// if get here the prefix was equal in every str
		prefixCount++
		lngCmnPrx = lastStrPrefix
	}

	return lngCmnPrx
}