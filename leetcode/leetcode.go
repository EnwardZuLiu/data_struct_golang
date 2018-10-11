package leetcode

import (
	"fmt"
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

// leetcode 20 有效的括号
func isValid(s string) bool {
	dict := map[rune]int{'(': -1, ')': 1, '[': -2, ']': 2, '{': -3, '}': 3}

	runes := []rune(s)
	if len(runes) == 0 {
		return true
	}

	stack := make([]rune, 10)
	stackTop := -1

	for i := 0; i < len(runes); i++ {
		if dict[runes[i]] < 0 {
			stackTop++
			stack[stackTop] = runes[i]
		} else {
			if stackTop < 0 { // stack is empty, is false
				return false
			}
			flag := stack[stackTop]
			stackTop--
			if dict[flag]+dict[runes[i]] == 0 {
				continue
			} else {
				return false
			}
		}
	}
	if stackTop >= 0 {
		return false
	}

	return true
}

// 删除排序数组中的重复项  26
func removeDuplicates(nums []int) int {
	size := len(nums)
	for i := 0; i < size-1; {
		if nums[i] == nums[i+1] {
			shiftElement(i, &nums)
			size = size - 1
		} else {
			i++
		}
	}
	for i := 0; i < size; i++ {
		fmt.Print(i)
	}
	return size
}

// 位移数组中的元素
func shiftElement(index int, nums *[]int) {
	size := len(*nums)
	if index < 0 || index >= size {
		return
	}
	for i := index; i < size-1; i++ {
		(*nums)[i] = (*nums)[i+1]
	}
}

// 双指针来实现该功能，由于数组是排序过的，所以 j 对应的元素只会越来越大，每次比 i 所指向的元素更大之后，给 i++赋值替换即可
func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var i int = 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
