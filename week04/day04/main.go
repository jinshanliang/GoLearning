package main

import "fmt"

func main() {
	var arr2 = []int{1, 2, 3, 4, 5}
	fmt.Printf("值是%v 长度是%v\n", arr2, len(arr2))

	var arr1 []int
	fmt.Println(arr1 == nil)

	//切片的循环遍历
	var strslice = []string{
		"java",
		"node.js",
		"python",
		"golang",
		"c++",
	}
	for i := 0; i < len(strslice); i++ {
		fmt.Printf("切片的第%v个内容是：%v\n", i+1, strslice[i])
	}

	for _, v := range strslice {
		fmt.Println("切片的内容是：\n", v)
	}

	//使用make函数创建一个切片
	var a = make([]int, 3, 3)
	fmt.Println(a)
}
