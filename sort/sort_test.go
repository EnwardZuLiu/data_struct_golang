package sort

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

// []int 表示切片为引用类型
// [10]int 表示数组类型，为值类型

func Test_SelectSort(t *testing.T) {
	var inputArr [10]int
	randSource := rand.NewSource(time.Now().Unix())
	r := rand.New(randSource)
	for i := 0; i < 10; i++ {
		inputArr[i] = r.Intn(10)
	}
	log.Println(inputArr)
	slice := SelectionSort(inputArr[:])
	log.Println(slice)
}

func Test_BubbleSort(t *testing.T) {
	var inputArr [10]int
	randSource := rand.NewSource(time.Now().Unix())
	r := rand.New(randSource)
	for i := 0; i < 10; i++ {
		inputArr[i] = r.Intn(10)
	}
	log.Println(inputArr)
	slice := BubbleSort(inputArr[:])
	log.Println(slice)
}

func Test_InsertSort(t *testing.T) {
	var inputArr [10]int
	randSource := rand.NewSource(time.Now().Unix())
	r := rand.New(randSource)
	for i := 0; i < 10; i++ {
		inputArr[i] = r.Intn(10)
	}
	log.Println(inputArr)
	slice := InsertSort(inputArr[:])
	log.Println(slice)
}

func Test_ShellSort(t *testing.T) {
	var inputArr [10]int
	randSource := rand.NewSource(time.Now().Unix())
	r := rand.New(randSource)
	for i := 0; i < 10; i++ {
		inputArr[i] = r.Intn(10)
	}
	log.Println(inputArr)
	slice := ShellSort(inputArr[:])
	log.Println(slice)
	log.Println("-----------")
}

func TestHeapSort(t *testing.T) {
	var inputArr [10]int
	randSource := rand.NewSource(time.Now().Unix())
	r := rand.New(randSource)
	for i := 0; i < 10; i++ {
		inputArr[i] = r.Intn(10)
	}
	log.Println(inputArr)
	HeapSort(inputArr[:])
	log.Println(inputArr)
}
