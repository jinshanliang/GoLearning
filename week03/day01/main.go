package main

import "fmt"

func main() {
	//打印 1- 100 之间所有是 9的倍数的整数的个数和总和
	/* 	var sum = 0
	   	var count = 0

	   	for i := 1; i <= 100; i++ {
	   		if i%9 == 0 {
	   			sum += i
	   			count++
	   		}
	   	}
	   	fmt.Printf("是九的倍数的和是：%v\n", sum)
	   	fmt.Printf("个数是%v\n", count) */

	//计算5的阶乘。1*2*3*4*5
	/* 	var sum = 1
	   	for i := 1; i < 6; i++ {
	   		sum *= i
	   	}
	   	fmt.Println(sum) */

	//打印矩形，使用for循环的嵌套
	/* 	var row = 3
	   	var column = 4

	   	for i := 0; i < row; i++ {
	   		for j := 0; j < column; j++ {
	   			fmt.Print("*")
	   		}
	   		fmt.Println("")
	   	} */

	//打印三角形
	/*
		var row = 5
		for i := 0; i < row; i++ {
			for j := 0; j < i; j++ {
				fmt.Print("*")
			}

			fmt.Println("")
		} */

	/* 	//打印九九乘法表

	   	for i := 1; i < 9; i++ {
	   		for j := 1; j <= i; j++ {
	   			fmt.Printf("%v*%v=%v \t", i, j, i*j)
	   		}

	   		fmt.Println("")
	   	} */

	//使用for range遍历字符串
	var str = "你好golang"
	for key, value := range str {
		fmt.Printf("key=%v,value=%c\n", key, value)
	}
}
