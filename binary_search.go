package main

import "fmt"

func binary_search(arr []int, low, high, hkey int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2

	if arr[mid] > hkey {
		return binary_search(arr, low, mid-1, hkey)
	} else if arr[mid] < hkey {
		return binary_search(arr, mid+1, high, hkey)
	}

	return mid
}
func main() {
	nums := make([]int, 0)
	for i := 0; i < 10; i++ {
		nums = append(nums, i)
	}
	fmt.Println(nums)
	fmt.Println(binary_search(nums, 0, len(nums), 8))
}
