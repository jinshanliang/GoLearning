package main

import "fmt"

func main() {
	var a float32 = 1.022
	var c float32 = 3.1415926535
	fmt.Printf("值：%v 类型：%T\n", a, a)
	fmt.Printf("值：%v 保留两位小数：%.4f", c, c)
}
