package main

import (
	"fmt"
	"sort"
)

func main() {
	a1 := []string{"tes", "ing", "owo"}
	a2 := []string{"test", "testing", "helloworld"}

	r := InArray(a1, a2)
	fmt.Println(r)
}

func AreIn(substr string, strings []string, rc chan<- string) {
	for _, str := range strings {
		j := 0
		for j < len(str) {
			x := 0
			for x+j < len(str) && x < len(substr) && str[x+j] == substr[x] {
				x++
			}
			if x == len(substr) {
				rc <- substr
				return
			}
			j++
		}
	}
}

func InArray(array1 []string, array2 []string) []string {
	r := make([]string, 0, len(array1))
	rc := make(chan string)

	i := 0
	for i < len(array1) {
		go AreIn(array1[i], array2, rc)
		i++

		substr := <-rc
		for _, s := range r {
			if s == substr {
				continue
			}
		}
		r = append(r, substr)
	}
	sort.Strings(r)
	return r
}
