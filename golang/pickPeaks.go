package main

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	peaks := PosPeaks{Pos: make([]int, 0), Peaks: make([]int, 0)}
	left := 0
	for left < len(array)-1 {
		right := left + 1
		i := right + 1
		for right < len(array) {
			if array[right] > array[left] {
				i = right + 1
				for i < len(array) {
					if array[right] > array[i] {
						peaks.Pos = append(peaks.Pos, right)
						peaks.Peaks = append(peaks.Peaks, array[right])
						break
					} else if array[right] != array[i] {
						right = i
					}
					i++
				}
				break
			} else {
				left++
				right++
			}
		}
		left = i
	}
	return peaks
}
