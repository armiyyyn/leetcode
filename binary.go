package main

import (
	"fmt"
)

func main() {
	nums := []int{2}
	var target int
	var res int = -1
	var mid int
	fmt.Println("give me a target:")
	fmt.Scan(&target)
	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[mid] == target {
			res = mid
			break
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = ((left + right)/2)
	}
	fmt.Println(res)
}
