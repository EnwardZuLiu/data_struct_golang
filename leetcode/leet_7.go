package leetcode

import (
	"math"
)

// 翻转整数
func reverse(x int) int {
	var rev int = 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32 && pop > 7) {
			return 0
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32 && pop < -8) {
			return 0
		}
		rev = pop + rev*10
	}
	return rev
}
