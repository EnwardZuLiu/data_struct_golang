package leetcode

func lengthOfLongestSubstring(s string) int {
	var hash = map[rune][]int{}
	runes := []rune(s)

	subStringLength := 0
	isRepeat := false
	for i := 0; i < len(runes); i++ {
		if value, ok := hash[runes[i]]; ok {
			isRepeat = true
			hash[runes[i]] = append(value, i)
			value, _ := hash[runes[i]]
			if flag := value[len(value)-1] - value[len(value)-2]; flag > subStringLength {
				subStringLength = flag
			}
			if value[0]+1 > subStringLength {
				subStringLength = value[0] + 1
			}
		} else {
			hash[runes[i]] = []int{i}
		}
	}
	if !isRepeat {
		return len(runes)
	}
	return subStringLength
}
