package main

import (
	"fmt"
	"log"
)

func main() {
	//创建一个切片
	s := []int{1, 2, 3, 4}
	fmt.Println(s)

	//使用make函数创建一个切片
	t := make([]string, 3, 6)
	t[0] = "java"
	t[1] = "go"
	t[2] = "python"
	fmt.Println(t)

	//创建一个包含字符串 "a","b","c" 的切片
	str := []string{"a", "b", "c"}
	fmt.Println(str)

	//创建一个长度为5的int切片（元素默认0）
	st := make([]int, 5)
	fmt.Printf("st的长度是：%v,容量是：%v\n", len(st), cap(st))
	//定义切片的长度是2 ，容量是10
	st3 := make([]int, 2, 10)
	fmt.Printf("st3的长度是：%v,容量是：%v\n", len(st3), cap(st3))
	//arr := [5]int{10, 20, 30, 40, 50}，在切片中取出【30，40】
	arr := [5]int{10, 20, 30, 40, 50}
	st4 := arr[2:4]
	fmt.Println(st4)

	//判断下面的语法是否正确，为什么？
	// A. s := []int{1,2,3}----对
	// B. s := make([]int) ---不对，未定义长度和容量，容量可以不定义，但是长度必须定义
	// C. s := make([]int, 5) -----对
	// D. s := make([]int, 5, 10) ---对
	// E. s := new([]int)----不对，没有new函数
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1))

	s2 := make([]int, 5)
	fmt.Println(len(s2), cap(s2))

	s4 := [3]int{10, 20, 30}
	s5 := s4[0:2]
	fmt.Println(len(s5), cap(s5))

	//创建一个长度为2、容量为10的切片，打印长度和容量
	arr1 := make([]int, 2, 10)
	fmt.Printf("arr1的长度是%v，容量是%v\n", len(arr1), cap(arr1))
	/* log.Println("这是一条很普通的日志,日志起到什么作用呢") */
}
