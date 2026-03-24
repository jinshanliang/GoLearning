package main

import (
	"fmt"
	"strconv"
)

func main() {

	str := "123456"
	//字符串转换成整型
	num, _ := strconv.ParseInt(str, 10, 64)
	fmt.Println(num)
}
