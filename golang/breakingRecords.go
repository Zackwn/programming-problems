package main

func breakingRecords(scores []int32) []int32 {
	var brokeBest int32 = 0
	var brokeWorst int32 = 0
	minScore := scores[0]
	maxScore := scores[0]

	for index := 1; index < len(scores); index++ {
		score := scores[index]
		if score > maxScore {
			maxScore = score
			brokeBest++
		} else if score < minScore {
			minScore = score
			brokeWorst++
		}
	}

	return []int32{brokeBest, brokeWorst}
}
