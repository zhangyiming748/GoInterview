package main

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

//27移除元素
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	var j int = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

//28. 实现 strStr()
//func strStr(haystack string, needle string) int {
//	if len(haystack)==0&&len(needle)==0{
//		return 0
//	}
//	rHay:=[]rune(haystack)
//	rNeed:=[]rune(haystack)
//	for i,v:=range rHay{
//		fmt.Printf("%v\t%v\n",i,string(v))
//		fmt.Printf("%v\n",string(rNeed[0:1]))
//		if
//	}
//	return 0
//}
//15. 三数之和
//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
func threeSum(nums []int) [][]int {
	answer := make([][]int, 0)
	var (
		a int
		b int
		c int
	)
	for i := 0; i < len(nums); i++ {
		a = 0 - nums[i]
		for j := i + 1; i < len(nums); j++ {
			b = nums[j]
			for k := j + 1; k < len(nums); j++ {
				c = nums[k]
				if b+c == a {
					enable := make([]int, 0)
					enable = append(enable, a, b, c)
					answer = append(answer, enable)
				}
			}
		}
	}
	return answer
}
func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}

	threeSum(nums)
}
