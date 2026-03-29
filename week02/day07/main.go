package main

import "fmt"

func main() {
	/*
		//1.最简单的if语句
		flag := true
		if flag {
			fmt.Println("这是一个真的语句")
		}

		age := 30
		if age > 18 {
			fmt.Println("成年人！")
		} */

	/* 	if age := 18; age >= 18 {
		fmt.Print("成年人了")
	} */

	//打印1-10的所有数据
	/* 	var i = 1
	   	for ; i < 11; i++ {
	   		fmt.Println(i)
	   	} */
	//打印1-50之间的偶数
	/* 	for i := 1; i < 51; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	} */
	//1+2+3.....+100的和
	sum := 0
	for i := 1; i <= 100; i++ {

		sum += i

	}
	fmt.Println(sum)
}
