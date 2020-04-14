//给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
//
//返回被除数 dividend 除以除数 divisor 得到的商。
//
//整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

package main

import (
	"fmt"
)

func main() {
	devid := 10
	divis := 3
	v := div(devid, divis)
	fmt.Println(v)
}

func getsAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
func isNegative(dividend, divisor int) bool {
	if dividend > 0 && divisor > 0 {
		return false
	} else if dividend < 0 && divisor < 0 {
		return false
	} else {
		return true
	}
}
func div(dividend, divisor int) int {
	var (
		absdend int = getsAbs(dividend)
		abssor  int = getsAbs(divisor)
		count   int = 0
		limit   int = absdend
	)
	for i := 0; i <= limit; i++ {
		if absdend >= abssor {
			count++
			absdend = absdend - abssor
		}
	}
	if isNegative(dividend, divisor) {
		count = -count
	}
	return count
}
