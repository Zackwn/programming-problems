package main

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	i := 0
	si := 0
	for i < len(array) && si < len(sequence) {
		if array[i] == sequence[si] {
			si++
		}
		i++
	}
	return si == len(sequence)
}
