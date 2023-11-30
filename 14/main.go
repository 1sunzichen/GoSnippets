package main

import "fmt"

func main() {
	nums := []int{2, 4, 1, 0, 3, 5}
	quickSortV2(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
func partition(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j-- //从右到左找出首个小于基准数的元素
		}
		for i < j && nums[i] <= nums[left] {
			i++ //从左到右找首个大于基准数的元素
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}
func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	p := partition(nums, left, right)
	quickSort(nums, left, p-1)
	quickSort(nums, p+1, right)
}

// 尾递归
func quickSortV2(nums []int, left, right int) {
	for left < right {
		pivot := partition(nums, left, right)

		if pivot-left < right-pivot {
			quickSortV2(nums, left, pivot-1)
			left = pivot + 1
		} else {
			quickSortV2(nums, pivot+1, right)
			right = pivot - 1
		}
	}
}
