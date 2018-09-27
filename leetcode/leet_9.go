package leetcode

// 判断是否是回文数字
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	var rev int = 0
	var temp int = x
	for temp != 0 {
		pop := temp % 10
		temp /= 10
		rev = pop + rev*10
	}

	if x == rev {
		return true
	}
	return false
}
