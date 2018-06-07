package util

import "math"

func Float2Int(x float64) (value int, error string) {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, faction := math.Modf(x)
		if faction >= 0.5 {
			whole++
		}
		return int(whole), "success"
	} else {
		return -1, "error"
	}

}
