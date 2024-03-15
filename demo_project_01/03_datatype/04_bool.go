package main

import "fmt"

func main() {
	//定义布尔类型的数据：
	var flag bool = true
	fmt.Println(flag)
	var flag1 bool = false
	fmt.Println(flag1)
	//布尔类型，底层是int类型，0表示false，1表示true
	var flag2 int = 1
	fmt.Println(flag2)

	//测试布尔类型的数值：
	var flag01 bool = true
	fmt.Println(flag01)
	var flag02 bool = false
	fmt.Println(flag02)
	var flag03 bool = 5 < 9
	fmt.Println(flag03)
}
