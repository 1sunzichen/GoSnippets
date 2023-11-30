package main

//插入排序
import "fmt"

func main() {
	n := []int{4, 3, 2, 1}
	fmt.Println(insertionSort(n))
}
func insertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		base := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > base {
			nums[j+1] = nums[j]
			j--
		}
		fmt.Println("i", j+1)
		nums[j+1] = base
	}
	return nums
}
