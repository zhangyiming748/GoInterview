package main

import "fmt"

//27.移除元素
func removeElement(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	i := 0
	j := 0
	for j := 0; j < length; j++ {
		if nums[j] == val {
			j++
		} else {
			nums[i] = nums[j]
			i++
			j++
		}
	}
	return length - (j - i)

}

//136.只出现一次的数字
func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res = res ^ v
	}
	return res
}

//1.两数之和
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		if i >= target {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//7.整数反转
func reverse(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10
		if !(-(1<<31) <= y && y <= (1<<31)-1) {
			return 0
		}
		x = x / 10
	}
	return y
}

//29.两数相除
func divide(dividend int, divisor int) int {
	isNegative := false
	if dividend < 0 {
		isNegative = !isNegative
		dividend = 0 - dividend
	}
	if divisor < 0 {
		isNegative = !isNegative
		divisor = 0 - divisor
	}
	if divisor == 1 {
		if isNegative {
			dividend = 0 - dividend
		}
		if dividend > (1<<31-1) || dividend < (0-1<<31) {
			return dividend
		}


	}
	count := 0
	for i := dividend; i >= divisor; i = i - divisor {
		count++
	}
	if isNegative {
		count = 0 - count
	}
	return count
}
func main() {
	var line []int = []int{1, 2, 1, 2, 3, 7}
	s := singleNumber(line)
	fmt.Println(s)
	t := twoSum(line, 9)
	fmt.Println(t)
	num := 1324
	n := reverse(num)
	fmt.Println(n)
	var (
		dend = -10
		sor  = 3
	)
	dev := divide(dend, sor)
	fmt.Println(dev)
}