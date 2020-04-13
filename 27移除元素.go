package main

func removeElement(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	i := 0
	j := 0
	for j := 0; j<length; j++ {
		if nums[j] == val{
			j++
		}else {
			nums[i] = nums[j]
			i++
			j++
		}
	}
	return length - (j - i)

}
