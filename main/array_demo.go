package main

import (
	"fmt"
)

/**
 * go 语言中，slice 本质是是对底层数据的一个应用。
 * slice 有时候可以看到整个数组，同时也可能看到的是数组的一部分。
 * 如果我们修改slice之后，这样我看其他slice引用的这个数组也会改变。
 */
func main() {
	s1 := make([]int, 5)
	fmt.Printf("The length of s1: %d \n", len(s1))
	fmt.Printf("The capacity of s1: %d \n", cap(s1))
	fmt.Printf("The value of s1: %d \n", s1)

	s2 := make([]int, 5, 8)
	fmt.Printf("The length of s2: %d \n", len(s2))
	fmt.Printf("The capacity of s2: %d \n", cap(s2))
	fmt.Printf("The value of s2: %d \n", s2)

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d \n", len(s4))
	fmt.Printf("The capacity of s4: %d \n", cap(s4))
	fmt.Printf("The value of s4: %d \n", s4)

	s4[0] = 10
	fmt.Printf("The value of s3: %d \n", s3)
	fmt.Printf("The value of s4: %d \n", s4)
}
