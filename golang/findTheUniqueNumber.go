package main

func FindUniq(arr []float32) float32 {
	var repeated float32
	if arr[0] == arr[1] || arr[0] == arr[2] {
		repeated = arr[0]
	} else if arr[1] == arr[2] {
		repeated = arr[1]
	}
	for _, n := range arr {
		if n != repeated {
			return n
		}
	}
	return 0
}
