package main

//  选择排序 最小的 放在最左边
import "fmt"

func main() {
	var num = []int{1, 3, 2, 4}
	fmt.Println(selectSort(num))
}
func selectSort(num []int) []int {
	n := len(num)
	for i := 0; i < n; i++ {
		k := i
		for j := i + 1; j < n; j++ {
			if num[j] < num[k] {
				k = j
			}
		}
		num[i], num[k] = num[k], num[i]
	}
	return num
}
