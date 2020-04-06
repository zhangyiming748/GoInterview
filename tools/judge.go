package tools

import (
"fmt"
"reflect"
)
func Typey(i interface{}){
	fmt.Println("类型为: ", reflect.TypeOf(i))
}
func main() {
	v1 := "123456"
	v2 := 12

	// reflect
	fmt.Println("v1 type:", reflect.TypeOf(v1))
	fmt.Println("v2 type:", reflect.TypeOf(v2))
	typey(v1)
}