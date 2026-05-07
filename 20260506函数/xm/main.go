package main

import "fmt"

// TODO 1: 写一个打招呼函数（练习单参数+无返回值）
// 函数名：greet
// 功能：接收一个名字，打印 "你好, [名字]"
func greet(name string) {
	fmt.Println("我的名字：", name)
}

// TODO 2: 写一个加法函数（练习多参数+有返回值）
// 函数名：add
// 功能：接收两个整数，返回它们的和
func add(a, b int) int {
	return a + b
}

// TODO 3: 写一个计算圆面积的函数（练习浮点数+返回值）
// 函数名：circleArea
// 参数：半径 radius (float64)
// 返回值：面积 (float64)  公式：π * r * r  (π取3.14)

func circleArea(radius float64) float64 {
	return 3.14 * radius * radius
}
func main() {
	//第一个
	greet("梁金山")
	he := add(1, 2)
	fmt.Println(he)
	s1 := circleArea(2.2)
	fmt.Println(s1)
}
