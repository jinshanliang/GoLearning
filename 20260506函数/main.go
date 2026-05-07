package main

import "fmt"

// 计算乘积
func multiply(x, y int) int {
	q := x * y
	return q
}

//判断一个数字是否是偶数

func isEven(num int) bool {
	return num%2 == 0

}

// 三个数取最大值
func maxOfThree(a, b, c int) int {
	maxnum := a
	if b > maxnum {
		maxnum = b
	} else if c > maxnum {
		maxnum = c
	}
	return maxnum
}

func main() {
	// 要求：写一个函数 multiply，接收两个整数，返回它们的乘积
	// 调用示例：multiply(4, 5) 应该返回 20
	result := multiply(4, 5)
	fmt.Println(result)

	// 要求：写一个函数 isEven，接收一个整数，返回是否是偶数
	// 偶数：能被2整除的数
	// 提示：用取模运算符 %，如 4%2 == 0 说明是偶数
	ress := isEven(5)
	fmt.Println(ress)

	// 要求：写一个函数 maxOfThree，接收三个整数，返回最大的那个
	res2 := maxOfThree(1, 90, 3)
	fmt.Printf("三个数的最大数字是%v \n", res2)
}
