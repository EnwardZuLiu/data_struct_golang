package leetcode

import (
	"math"
)

//leet 14， 最长公共前缀
func longestCommonPrefix(strs []string) string {
	strSize := len(strs)
	if strSize == 0 {
		return ""
	}
	if strSize == 1 {
		return strs[0]
	}
	runes := make([][]rune, strSize)
	var length int = math.MaxInt32
	for i := 0; i < strSize; i++ {
		runes[i] = []rune(strs[i])
		if length > len(runes[i]) {
			length = len(runes[i])
		}
	}

	var prefixNumber int = -1
	var breakFlag bool = false

	for i := 0; i < length; i++ {
		if breakFlag {
			break
		}

		flag := runes[0][i]
		for j := 0; j < len(runes); j++ {
			if flag != runes[j][i] {
				prefixNumber = i
				breakFlag = true
				break
			}
		}
	}
	if !breakFlag {
		prefixNumber = length
	}
	if prefixNumber == -1 {
		return ""
	}
	return string(runes[0][0:prefixNumber])
}
