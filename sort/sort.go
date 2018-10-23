package sort

/**
 * 时间复杂度的定义： 循环执行数量次数最多的那一段代码（不代表循环本身）
 */

/**
 * 选择排序算法，每次在未排序的序列中找到最小（最大）的元素，存放到排序位置的起始位置，
 * 然后在从剩下的未排序的序列中找到最小（最大）的元素，存放到已排序的序列的后面
 *  最好时间复杂度： O(n^2)
 *  最坏时间复杂度: O(n^2)
 *  平均时间复杂度： O(n^2)
 *  均摊事件复杂度：
 *  不稳定排序算法
 */
func SelectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

/**
 * 冒泡排序：它重复地走访过要排序的数列，一次比较两个元素，如果他们的顺序错误就把他们交换过来。走访数列的工作是重复地进行直到没有再需要交换，
 * 也就是说该数列已经排序完成。这个算法的名字由来是因为越小的元素会经由交换慢慢“浮”到数列的顶端。
 *  最好时间复杂度： O(n^2)
 *  最坏时间复杂度: O(n^2)
 *  平均时间复杂度： O(n^2)
 */
func BubbleSort(arr []int) []int {
	size := len(arr)
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

/**
 * 插入排序：它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入
 *  最好时间复杂度： O(n)
 *  最坏时间复杂度: O(n^2)
 *  平均时间复杂度： O(n)
 *  稳定排序算法
 */
func InsertSort(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
	return arr
}

func ShellSort(arr []int) []int {
	length := len(arr)
	h := 1
	for h < length/3 {
		h = h*3 + 1
	}
	for h >= 1 {
		for i := h; i < length; i++ {
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				arr[j], arr[j-h] = arr[j-h], arr[j]
			}
		}
		h /= 3
	}
	return arr
}

func HeapSort(arr []int) {
	N := len(arr) - 1
	for k := N / 2; k >= 1; k-- {
		sink(arr, k, N)
	}
	for N >= 1 {
		arr[1], arr[N] = arr[N], arr[1]
		N--
		sink(arr, 1, N)
	}
	for j := 0; j < len(arr)-1; j++ {
		if arr[j] > arr[j+1] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}

func sink(s []int, k, N int) {
	for {
		i := 2 * k
		if i > N {
			break
		}
		if i < N && s[i+1] > s[i] {
			i++
		}
		if s[k] > s[i] {
			break
		}
		s[k], s[i] = s[i], s[k]
		k = i
	}
}
