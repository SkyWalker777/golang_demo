package main

import "fmt"

/*
*	变量的使用方式
 */
var name2, age2, sex2, heigh2 = "wangwu", 40, "男", 178.9

var (
	name3  = "huangfu"
	age3   = 30
	sex3   = "man"
	heigh3 = 178.9
)

func main() {
	// 第一种：在声明变量时赋值
	var num int = 10
	fmt.Println(num)

	// 第二种：先声明变量，再赋值
	var num1 int
	num1 = 12
	fmt.Println(num1)

	// 第三种：只声明，不赋值，使用类型的默认值
	var num2 int
	fmt.Println(num2) // 输出 0

	// 第四种：不声明类型，直接给变量赋值，会根据赋值内容判定该变量类型（自动类型推断）
	var num3 = 14
	var str = "abc"
	fmt.Println(num3)
	fmt.Println(str)

	// 第五种：使用 := 忽略var关键词
	str2 := "zhangsan"
	fmt.Println(str2)

	// 第六种：多变量一起声明
	var num4, num5, num6 int
	fmt.Println(num4, num5, num6) // 输出 0 0 0

	var (
		num7 = 20
		num8 = 21
		num9 = 22
	)
	fmt.Println(num7, num8, num9) // 输出 20 21 22

	var name, age, sex, heigh = "lisi", 20, "男", 17.8
	fmt.Println(name, age, sex, heigh)                  // 输出 lisi 20 男 17.8
	fmt.Printf("%T %T %T %T \n", name, age, sex, heigh) // string int string float64

	name1, age1, sex1, heigh1 := "lisi", 20, "男", 17.8
	fmt.Println(name1, age1, sex1, heigh1) // 输出 lisi 20 男 17.8

	// 第七种：全局变量 定时在代码块{}外部
	fmt.Println(name2, age2, sex2, heigh2) // 输出 wangwu 40 男 178.9
	var name2, age2, sex2, heigh2 = "wangwu2", 42, "女", 190.4
	fmt.Println(name2, age2, sex2, heigh2) // wangwu2 42 女 190.4

	fmt.Println(name3, age3, sex3, heigh3) // huangfu 30 man 178.9
}
