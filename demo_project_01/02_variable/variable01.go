package main

import "fmt"

func main() {
	//1.变量的声明
	var age int
	//2.变量的赋值
	age = 18
	//3.变量的使用
	fmt.Println("age = ", age) /*输出：age =  18*/
	//声明和赋值可以合成一句：
	var age2 int = 19
	fmt.Println("age2 = ", age2) /*输出：age2 =  19*/
	/*变量的重复定义会报错：
	  ./variable01.go:19:6: age redeclared in this block
	        ./variable01.go:7:6: other declaration of age
	*/
	//var age int = 20
	//fmt.Println("age = ", age)

	//不可以在赋值的时候给与不匹配的类型
	/**
	./variable01.go:23:16:
	cannot use 12.56 (untyped float constant) as int value in variable declaration (truncated)
	*/
	//var num int = 12.56
	//fmt.Println("num = ", num)
}
