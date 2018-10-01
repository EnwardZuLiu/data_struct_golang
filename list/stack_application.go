package list

import (
	"fmt"
)

/**
 * 输入一个非负数的十进制，打印其 n 进制的数字
 */
func conversion(number uint32, n uint32) {
	stack := Stack{nil, 0}
	// 入栈
	for number != 0 {
		stack.Push(number % n)
		number = number / n
	}
	// 逐步打印所有的数据
	stack.Traverse(func(value interface{}) {
		v := fmt.Sprintf("%v", value) // 将 interface 转化为 string 类型
		fmt.Printf(v)
	})
}
