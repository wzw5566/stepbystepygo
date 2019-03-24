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
	route.POST("/postdata", postData)
	route.Run(":8085")
}
//BindJSON获取post 请求参数
func postData(c *gin.Context)  {
	var person Person

		if c.BindJSON(&person) == nil {
		log.Println("====== Bind By JSON ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "成功获取Post数据")

}
