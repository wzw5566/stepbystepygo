package main

import (
	"fmt"
	"stepbystepygo/chapter10/model"
)
var i = 10

func main() {
	//创建student的实例
	var stu = model.Student{
		Name:  "Vincent",
		Score: 78.87,
	}

	fmt.Println(stu)

}

