package leetcode

import (
	"testing"
)

func Test_leet13(t *testing.T) {
	sum := romanToInt("MCMXCIV")
	if sum == 1994 {
		t.Log("Pass")
	} else {
		t.Error("fail")
	}
}

func Test_leet14(t *testing.T) {
	str := longestCommonPrefix([]string{"dog", "racecar", "car"})
	if str == "" {
		t.Log("PASS")
	} else {
		t.Error("FAIL")
	}
	str1 := longestCommonPrefix([]string{"flower", "flow", "flight"})
	if str1 == "fl" {
		t.Log("PASS")
	} else {
		t.Error("FAIL")
	}
	str2 := longestCommonPrefix([]string{"aa", "a"})
	if str2 == "a" {
		t.Log("PASS")
	} else {
		t.Error("FAIL11" + str2)
	}

}
