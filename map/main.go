package main
import "fmt"

var age = test()
func test() int {
	fmt.Println("最先执行全局变量初始哈")
	return 90
}

func init() {

	fmt.Println("再执行init初始化函数")

}

func main() {
	fmt.Println("最后执行main函数")

}
