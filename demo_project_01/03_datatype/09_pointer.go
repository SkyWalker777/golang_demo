package main

import "fmt"

/*
指针类型
*/
func main() {
	var age int = 18
	// &+变量 获取变量的内存的地址，
	fmt.Println(&age) // 0x1400001e088

	//定义一个指针变量：
	//var代表要声明一个变量
	//ptr 指针变量的名字
	//ptr对应的类型是：*int 是一个指针类型 （可以理解为 指向int类型的指针）
	//&age就是一个地址，是ptr变量的具体的值
	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr本身这个存储空间的地址为：", &ptr)
	//想获取ptr这个指针或者这个地址指向的那个数据：
	fmt.Printf("ptr指向的数值为：%v \n", *ptr) //ptr指向的数值为：18

	// *+内存地址可以获取到内存地址指向的变量内容 可以理解为内存地址就是一个指针变量
	fmt.Printf("ptr指向的数值为：%v \n", *&age) // 18

	// 可以通过指针改变指向的变量值
	*ptr = 20
	fmt.Printf("ptr指向的新数值为：%v \n", *ptr) //ptr指向的新数值为：18

}
