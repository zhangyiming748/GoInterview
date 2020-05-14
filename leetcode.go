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
func singleNumberstand(nums []int) int {
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

//69. x 的平方根
//使用二分法查找最接近的整数答案
func mySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)>>1
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

//3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	for i := range s {

	}
}

type times int

//136
func singleNumber(nums []int) int {
	once := make(map[int]times)
	for _, v := range nums {

		once[v]++

	}
	for key, value := range once {
		if value == 1 {
			return key
		}
	}

	return 0
}

//287
func findDuplicate(nums []int) int {
	once := make(map[int]bool)
	for _, v := range nums {
		if _, ok := once[v]; ok {
			return v
		} else {
			once[v] = true
		}
	}
	return 0
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
