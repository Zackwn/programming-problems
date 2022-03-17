package kata

import "strconv"

func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}
	catQuantity := make([]int, len(listCat))
	for _, stock := range listArt {
		for i, cat := range listCat {
			if string(stock[0]) == cat {
				s, e := len(stock)-1, len(stock)
				for stock[s] != ' ' {
					s--
				}
				s++
				q, _ := strconv.Atoi(stock[s:e])
				catQuantity[i] += q
			}
		}
	}
	var result string
	for i, cat := range listCat {
		if i != 0 {
			result += " - "
		}
		result += "(" + cat + " : " + strconv.Itoa(catQuantity[i]) + ")"
	}
	return result
}
