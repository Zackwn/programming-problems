package main

func allLongestStrings(inputArray []string) []string {
	strmap := make(map[int][]string)
	longest := 0
	for _, str := range inputArray {
		l := len(str)
		if l > longest {
			longest = l
		}
		strmap[l] = append(strmap[l], str)
	}
	return strmap[longest]
}
