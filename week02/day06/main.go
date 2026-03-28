package main

import "fmt"

func main() {

	//第一题：有两个变量，a和b，要求将其转换，并打印出来
	// var a = 10
	// var b = 20
	// var c int
	// c = a
	// a = b
	// b = c
	// fmt.Println(a, b)

	// //第二题：
	// var a = 10
	// var b = 20
	// b, a = a, b
	// fmt.Print("a=%v b=%v", a, b)

	//100天以后还有多少周多少天
	var week = 100 / 7
	var days = 100 % 7
	fmt.Printf("距离放假还有%v周%v天", week, days)
}
