package main

import (
	"fmt"
)
/*
		 type 结构体名称 struct
		 		

*/

// 定义结构体
type cat struct {
	Name  string
	Age   int
	color string
}

func main() {

	//创建结构体实例
	var cat1 cat
	cat1.Name = "小白"

	var cat2 cat
	cat2.Name = "小花"
	fmt.Println("第一只猫叫：", cat1.Name)
	fmt.Println("第二只猫名叫：", cat2.Name)

}
