package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// 整数边界演示
	/*
		cannot use -300 (untyped int constant) as int8 value in variable declaration (overflows)
		cannot use 300 (untyped int constant) as int8 value in variable declaration (overflows)
	*/
	//var num, num1 int8 = -300, 300
	//fmt.Println(num, num1)

	// 无符号整数边界演示
	//var num2 uint8 = -12 // cannot use -12 (untyped int constant) as uint8 value in variable declaration (overflows)
	//fmt.Println(num1)

	var num3 = 10
	fmt.Println("num3 =", num3)
	fmt.Printf("num3的类型是：%T \n", num3) // num3的类型是：int32  整形默认的类型是int
	fmt.Println(unsafe.Sizeof(num3))   // 变量num3占用的字节数
}
