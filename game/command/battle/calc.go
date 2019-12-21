package battle

import "math/rand"

func disperse(src, rangePercent int) int {
	min := 100 - rangePercent/2
	percent := rand.Intn(rangePercent) + min
	out := src * percent / 100
	if out <= 0 {
		return 1
	}
	return out
}
