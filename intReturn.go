package main
//整数反转
import (
	"fmt"
	"./tools"
	"strconv"
)
func reverse(x int) int {
	//是否为负数
	strx:=strconv.Itoa(x)
	fmt.Println(strx)
	for i,ch:=range strx{
		fmt.Printf("%d%d\n",i,ch)
	}
	fmt.Printf("%f",strx[2])
	//isNegative:=false
	if x<0{
		x=-x
		fmt.Println(x)
		//isNegative=true
	}
	//转换字符串
	slicex:=strconv.Itoa(x)
	tools.Typey(slicex)
	for i,v:=range slicex{
		fmt.Printf("第%v个数字是%v\t",i,v)
	}

	return x
}
func main()  {
	num1:=12345
	reverse(num1)
	//num2:=-123
}
