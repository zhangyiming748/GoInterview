package main
//整数反转
import (
	"fmt"
	"./tools"
)
func reverse(x int) int {
	//是否为负数

	//isNegative:=false
	if x<0{
		x=-x
		fmt.Println(x)
		//isNegative=true
	}
	//转换字符串
	slicex:=string(x)
	tools.Typey(slicex)

	return x
}
func main()  {
	num1:=12345
	reverse(num1)
	//num2:=-123
}
