package leetcode

// 罗马数字的翻转
func romanToInt(s string) int {
	dict := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	var sum int

	romanNumbers := []rune(s)
	size := len(romanNumbers)
	for i := 0; i < size; i++ {
		if i >= size-1 { // next is not exist
			sum += dict[romanNumbers[i]]
			continue
		}
		if dict[romanNumbers[i]] < dict[romanNumbers[i+1]] {
			sum += dict[romanNumbers[i+1]] - dict[romanNumbers[i]]
			i++
			continue
		}
		sum += dict[romanNumbers[i]]
	}

	return sum
}
