package sort

func SelectSort(arr []int) []int {
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

func InsertSort(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
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
	for k := N/2 ; k >= 1; k-- {
		sink(arr, k, N)
	}
	for N >= 1 {
		arr[1], arr[N] = arr[N], arr[1]
		N--
		sink(arr, 1 ,N)
	}
	for j := 0; j < len(arr) - 1; j++ {
		if arr[j] > arr[j+1] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}

func sink(s []int, k, N int) {
	for{
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

