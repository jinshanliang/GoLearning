package main

import "fmt"

func main() {

	var username string = "张三"
	var age = 10
	fmt.Println(username)
	// fmt.Print("张三的年龄是:", age)
	fmt.Printf("张三的年龄是：%d\n", age)
}
