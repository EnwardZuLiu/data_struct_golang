package leetcode

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	var nums = []int{1, 2, 3, 5, 6, 7, 8, 9, 10}
	target := 7
	result := twoSum(nums, target)
	fmt.Println(result)
}

func Test_twoSum2(t *testing.T) {
	var nums = []int{1, 2, 3, 5, 6, 7, 8, 9, 10}
	target := 7
	result := twoSum2(nums, target)
	fmt.Println(result)
}
