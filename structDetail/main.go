package main

import (
	"encoding/json"
	"fmt"
)

//结构体的初始化
type Person struct {
	Name string `json:"name"` //json 结构体标签
	Age  int    `json:"age"`
}

func main() {

	p2 := Person{"vincent", 33}
	jsonStr, err := json.Marshal(p2)
	if err != nil {
		fmt.Println("json处理错误,", err)
	}
	fmt.Println("jsonStr", string(jsonStr))
}
