package main

import (
	"fmt"
	"strings"
	_ "strings"
)

func main() {
	//定义字符串
	// var name1 string = "你好，golong"
	// var name2 = "你好，go！"
	// name3 := "你好，gogogo"
	// fmt.Println(name1,name2,name3)
	// var str1 = `
	// hello
	// this is my book
	// you lool a book!

	// `
	// fmt.Println(str1)

	// str1 := "你好"
	// str2 := "golang"

	// str3 := fmt.Sprintf("%v %v ", str1, str2)
	// fmt.Print(str3)

	// var num = "123-456-789"
	// str1 := strings.Split(num, "-")
	// fmt.Println(str1)
	// str2 := strings.Join(str1, "`")
	// fmt.Println(str2)

	// url := "https://golang.org/pkg"
	// /* 	这个字符串是否包含 "golang"？
	// 这个字符串是否以 "https://" 开头？
	// 这个字符串是否以 "/pkg" 结尾？ */
	// flag := strings.Contains(url, "golang")
	// fmt.Println(flag)
	// flag1 := strings.HasPrefix(url, "https://")
	// flag2 := strings.HasSuffix(url, "/pkg")
	// fmt.Println(flag1, flag2)

	data := "apple,banana,orange,grape"
	/*
	   查找 "banana" 在字符串中的起始位置（下标）
	   查找 "watermelon" 在字符串中的位置，观察返回值
	   用 strings.Split 将字符串按逗号分割成切片
	   将分割后的切片，用 " | " 连接成新的字符串 */
	one := strings.Index(data, "banana")
	fmt.Println(one)
	two := strings.Index(data, "watermelon")
	fmt.Println(two)

	three := strings.Split(data, ",")
	fmt.Println(three)

	four := strings.Join(three, "|")
	fmt.Println(four)
}
