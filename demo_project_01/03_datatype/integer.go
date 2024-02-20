package main

import "fmt"

func main() {

	// 整数边界演示
	/*
		cannot use -300 (untyped int constant) as int8 value in variable declaration (overflows)
		cannot use 300 (untyped int constant) as int8 value in variable declaration (overflows)
	*/
	var num, num1 int8 = -300, 300
	fmt.Println(num, num1)

	// 无符号整数边界演示
	var num2 uint8 = -12 // cannot use -12 (untyped int constant) as uint8 value in variable declaration (overflows)
	fmt.Println(num1)
}
