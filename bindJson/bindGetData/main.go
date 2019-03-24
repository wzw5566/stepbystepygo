package main

import "log"
import "github.com/gin-gonic/gin"

//定义接收请求的数据结构体
type Person struct {
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}

func main() {
	route := gin.Default()
	route.GET("/getdata", getData)
	route.Run(":8085")
}
//获取url中的请求参数
func getData(c *gin.Context) {
	var person Person
	if c.Bind(&person) == nil {
		log.Println("====== Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
//获取请求body中的data数据
	if c.BindJSON(&person) == nil {
		log.Println("====== Bind By JSON ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}

	c.String(200, "Success")
}