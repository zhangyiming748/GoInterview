package main

import (
	"strconv"
	"fmt"
	"./tools"
)

func isPalindrome(x int) bool {
	s:=strconv.Itoa(x)
	fmt.Printf("%v",s)
	tools.Typey(s)
	for i,ch:=range s{
		fmt.Printf("%v%v",i,ch)
	}
	return false
}
func main(){
	num:=121
	isPalindrome(num)
}