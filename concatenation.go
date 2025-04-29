package main

import "fmt" 

func main() {
	nums := []int{0}
	var res []int
	for i := 0; i < len(nums); i++ {
		res = append(res, nums[i])
		if i == len(nums)-1 {
			i -= len(nums)
		}
		if len(res) == len(nums)*2 {
			break
		}
	}
	return res
}
