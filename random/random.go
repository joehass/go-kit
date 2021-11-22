package random

import "math/rand"

// GetRandom 生成指定区间的随机数
func GetRandom(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}