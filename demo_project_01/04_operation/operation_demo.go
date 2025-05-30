package main

import "fmt"

func main() {
	//arithmeticOp()
	//switchOp1()
	switchOp2()
}

/*
算术运算符
*/
func arithmeticOp() {
	//+加号：
	//1.正数 2.相加操作  3.字符串拼接
	var n1 int = +10
	fmt.Println(n1)
	var n2 int = 4 + 7
	fmt.Println(n2)
	var s1 string = "abc" + "def"
	fmt.Println(s1)
	// /除号：
	fmt.Println(10 / 3)   //两个int类型数据运算，结果一定为整数类型
	fmt.Println(10.0 / 3) //浮点类型参与运算，结果为浮点类型
	// % 取模  等价公式： a%b=a-a/b*b
	fmt.Println(10 % 3) // 10%3= 10-10/3*3 = 1
	fmt.Println(-10 % 3)
	fmt.Println(10 % -3)
	fmt.Println(-10 % -3)
	//++自增操作：
	var a int = 10
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)
	//++ 自增 加1操作，--自减，减1操作
	//go语言里，++，--操作非常简单，只能单独使用，不能参与到运算中去
	//go语言里，++，--只能在变量的后面，不能写在变量的前面 --a  ++a  错误写法
}

/*
赋值运算符
*/
func assignmentOp() {
	var num1 int = 10
	fmt.Println(num1)
	var num2 int = (10+20)%3 + 3 - 7 //=右侧的值运算清楚后，再赋值给=的左侧
	fmt.Println(num2)
	var num3 int = 10
	num3 += 20 //等价num3 = num3 + 20;
	fmt.Println(num3)
}

/*
两值交换，引入一个中间变量实现
*/
func switchOp1() {
	var a int = 8
	var b int = 4
	fmt.Printf("原值：a = %v,b = %v \n", a, b)
	//交换：
	//引入一个中间变量：
	var t int
	t = a
	a = b
	b = t
	fmt.Printf("交换后：a = %v,b = %v", a, b)
}

/*
两值交换，直接替换
*/
func switchOp2() {
	var a int = 8
	var b int = 4
	fmt.Printf("原值：a = %v,b = %v \n", a, b)
	// 赋值时替换
	a, b = b, a
	fmt.Printf("交换后：a = %v,b = %v", a, b)
}

/*
关系运算符
*/
func compareOp() {

}

/*
逻辑运算
*/
func logicOp() {
	//与逻辑：&& :两个数值/表达式只要有一侧是false，结果一定为false
	//也叫短路与：只要第一个数值/表达式的结果是false，那么后面的表达式等就不用运算了，直接结果就是false  -->提高运算效率
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)
	//或逻辑：||:两个数值/表达式只要有一侧是true，结果一定为true
	//也叫短路或：只要第一个数值/表达式的结果是true，那么后面的表达式等就不用运算了，直接结果就是true -->提高运算效率
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)
	//非逻辑：取相反的结果：
	fmt.Println(!true)
	fmt.Println(!false)
}
