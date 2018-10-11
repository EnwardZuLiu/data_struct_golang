package leetcode

import (
	"fmt"
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

func Test_leet20(t *testing.T) {
	bool1 := isValid("")
	if !bool1 {
		t.Error("empty string is return ture")
	}

	bool2 := isValid("{}}")
	if bool2 {
		t.Error("{}} string is return false")
	}

	bool3 := isValid("{]")
	if bool3 {
		t.Error("{}} string is return false")
	}

	bool4 := isValid("{([])]")
	if bool4 {
		t.Error("{}} string is return false")
	}

	bool5 := isValid("{([])")
	if bool5 {
		t.Error("{}} string is return false")
	}

	bool6 := isValid("()[]{}")
	if !bool6 {
		t.Error("()[]{} string is return false")
	}
}

func Test_leet_21(t *testing.T) {
	node1 := ListNode{Val: 5}
	node2 := ListNode{Val: 6}
	node3 := ListNode{Val: 7}
	node4 := ListNode{Val: 2}
	node5 := ListNode{Val: 4}
	node6 := ListNode{Val: 6}
	node7 := ListNode{Val: 7}
	node8 := ListNode{Val: 9}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4

	node5.Next = &node6
	node6.Next = &node7
	node7.Next = &node8

	list := mergeTwoLists(&node1, &node5)
	fmt.Println(list)
}

func Test_for(t *testing.T) {
	for i := 0; i < 3; i++ {
		fmt.Print(i)
		i := i // The left i is a new declared variable,
		// fmt.Print(i)
		// and the right i is the loop variable.
		i = 10 // The new delcared variable is modified,
		fmt.Print(i)
		// but the old one (the loop variable) is not.
	}
}

func Test_removeDuplicates(t *testing.T) {
	i := removeDuplicates([]int{0, 0, 1, 1, 1, 2})
	if i == 3 {
		t.Log("SUCESSS")
	} else {
		t.Error("result should is 3")
	}
}

func Test_removeDuplicates1(t *testing.T) {
	i := removeDuplicates1([]int{0, 0, 1, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	if i == 10 {
		t.Log("SUCCESS")
	} else {
		t.Error("result should is 3")
	}
}
