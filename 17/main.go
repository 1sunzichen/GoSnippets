package main

import "fmt"

func siftDown(nums *[]int, n, i int) {
	for true {
		l := 2*i + 1
		r := 2*i + 2
		ma := i
		if l < n && (*nums)[l] > (*nums)[ma] {
			ma = l
		}
		if r < n && (*nums)[r] > (*nums)[ma] {
			ma = r
		}
		if ma == i {
			break
		}
		(*nums)[i], (*nums)[ma] = (*nums)[ma], (*nums)[i]
		i = ma
	}
}
func heapSort(nums *[]int) {
	//建成小顶堆 最大的在节点处
	for i := len(*nums)/2 - 1; i >= 0; i-- {
		siftDown(nums, len(*nums), i)
	}
	fmt.Printf("nums:%v\n", nums)
	for i := len(*nums) - 1; i > 0; i-- {
		(*nums)[0], (*nums)[i] = (*nums)[i], (*nums)[0]
		siftDown(nums, i, 0)
	}
}
func main() {
	arr := []int{41, 33, 51, 12, 22, 33, 60, 0}
	heapSort(&arr)
	fmt.Printf("arr:%v", arr)
}
