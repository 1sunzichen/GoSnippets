package main

import "fmt"

// 归并排序 真TM 难
func merge(nums []int, left, mid, right int) {
	tmp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}
	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}
	for j <= right {
		tmp[k] = nums[j]
		j++
		k++
	}
	for k := 0; k < len(tmp); k++ {
		nums[left+k] = tmp[k]
	}
}

func mergeSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)
	merge(nums, left, mid, right)
}
func main() {
	arr := []int{1, 2, 3, 5, 4, 6, 2}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr, "arr")
}
