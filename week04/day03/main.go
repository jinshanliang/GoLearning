package main

import "fmt"

func main() {

	var str1 = [3]int{1, 2, 3}
	fmt.Println(str1)

	var str2 = [4]string{"php", "golang", "python"}

	str2[0] = "phper"
	fmt.Println(str2)

	var str3 = [4]string{"php", "golang", "python"}
	for i := 0; i < len(str3); i++ {
		fmt.Println(str3[i])
	}

	for _, strkey := range str3 {
		fmt.Println(strkey)
	}
	//求一个数组的所有的和以及平均值，使用for 和 for range 来实现
	var str4 = [...]int{1, 2, 3, 4, 5, 6, 7}
	var sum = 0
	var ave = 0
	for i := 0; i < len(str4); i++ {
		sum += str4[i]
		ave = sum / len(str4)
	}
	fmt.Printf("数组里面的和：%v", sum)
	fmt.Printf("数组的和是：%v", ave)

	//判断一个数组的最大值，并求出对应的坐标
	var str5 = [...]int{1, 2, 3, 4, 5, 6, 7}
	max := 0
	index := 0
	for i := 0; i < len(str5); i++ {
		if str5[i] > max {
			max = str5[i]
			index = i
		}
	}
	fmt.Printf("数组的最大值是 %v 对应的索引下标是%v\n", max, index)

	//求一个数组中找出和为8的两个数的下标
	var str6 = [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(str6); i++ {
		for j := i + 1; j < len(str6); j++ {
			if str6[i]+str6[j] == 8 {
				fmt.Printf("(%v,%v)\n", i, j)
			}
		}
	}

	//定义一个二维数组

	var arr1 = [3][2]string{
		{"开发", "测试"},
		{"采购", "业务"},
		{"生产", "销售"},
	}
	fmt.Printf("获取到第一个的第一个:%v\n", arr1[0][0])
	//循环打印二维数组
	for _, v := range arr1 {
		for _, o := range v {
			fmt.Println(o)
		}
	}
}
