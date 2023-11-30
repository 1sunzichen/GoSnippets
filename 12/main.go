package main

// 冒泡排序 最大放在右边依次排序
import "fmt"

func main() {
	n := []int{2, 1, 4, 3}
	fmt.Println(bubbleSortWithFlag(n))
}

func bubbleSortWithFlag(nums []int) []int {
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
