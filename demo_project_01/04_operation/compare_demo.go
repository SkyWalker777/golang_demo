package main

import "fmt"

func main() {
	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{Name: "Alice", Age: 30}
	fmt.Println(p1 == p2) // 输出: true (内容相同)
	p1 = p2
	pp1 := &p1
	pp2 := &p2
	fmt.Println(pp1 == pp2) // 输出: false (即使内容相同，但它们是不同的实例)
	fmt.Printf("&p1: %v &p2: %v", &p1, &p2)

}

type Person struct {
	Name string
	Age  int
}
