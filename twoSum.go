package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	result := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result[0] = i
				result[1] = j
				//result = append(result, i)
				//result = append(result, j)
				return result
			}
		}
	}
	return result
}
func main() {
	nums := make([]int, 0)
	target := 9
	nums = append(nums, 2, 7, 11, 15)
	fmt.Println(twoSum(nums, target))
}
