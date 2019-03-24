package main

import (
	"fmt"
)

type Person struct {
	Name string
}

//给Person类型绑定一个方法test
func (p Person) test() {
	fmt.Println("name is", p.Name)
}

func main() {

	var p Person
	p.Name = "vincent~"

	p.test()// 方法调用

}
