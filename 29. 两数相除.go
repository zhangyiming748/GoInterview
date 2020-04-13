//给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
//
//返回被除数 dividend 除以除数 divisor 得到的商。
//
//整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

package main

import (
	"fmt"
)

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return 0
	}
	if dividend == divisor {
		return 1
	}

	count := 1
	for i := 0; i < dividend; i++ {
		if dividend-divisor >= divisor {
			dividend = dividend - divisor
			count++
		} else {
			break
		}
	}
	if isNegative(dividend, divisor) {
		count=-count
	}
	return count
}
func isNegative(dividend int, divisor int) bool {
	if dividend > 0 || divisor > 0 {
		return false
	} else if dividend < 0 || divisor < 0 {
		return false
	} else {
		return true
	}
}

func main() {
	devid := -4
	divis := 2
	v := divide(devid, divis)
	fmt.Println(v)
}
