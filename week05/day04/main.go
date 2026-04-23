package main

import (
	"fmt"
)

func main() {

	/* 	s := []int{5, 10, 15, 20, 25, 30, 35}
	   	// 截取 [15,20,25]
	   	s1 := s[2:5]
	   	fmt.Println(s1)
	   	// 截取 [5,10,15]
	   	s2 := s[0:3]
	   	fmt.Println(s2)
	   	// 截取 [25,30,35]
	   	s3 := s[4:]
	   	fmt.Println(s3)
	   	// 截取全部元素
	   	s4 := s[0:]
	   	fmt.Println(s4)

	   	//写出下列切片的长度和容量

	   	ss := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	   	aa := ss[2:5]
	   	fmt.Println(len(aa), cap(aa))
	   	b := ss[:4]
	   	fmt.Println(len(b), cap(b))

	   	c := ss[6:]
	   	fmt.Println(len(c), cap(c))

	   	d := ss[3:8]
	   	fmt.Println(len(d), cap(d)) */

	//将下面每个元素的值改为0
	/* 	s := []int{5, 10, 15, 20, 25}
	   	s[0] = 0
	   	s[1] = 0
	   	s[2] = 0
	   	s[3] = 0
	   	s[4] = 0
	   	fmt.Println(s) */

	//写代码打印每个元素
	/* 	nums := []int{2, 4, 6, 8, 10}
	   	for _, value := range nums {
	   		fmt.Println(value) */

	//使用for循环将下面的切片中的元素都乘以2
	nums := []int{1, 2, 3, 4, 5}
	nums2 := make([]int, 0, 5)
	for i := 0; i < len(nums); i++ {
		value := nums[i] * 2
		nums2 = append(nums2, value)
	}
	fmt.Println(nums2)
	//复制
	src := []int{10, 20, 30, 40, 50}
	// 创建dst，复制src的所有元素
	dst := append([]int(nil), src...)
	fmt.Println(dst)

	dst1 := make([]int, 3)
	copy(dst1, src)
	fmt.Println(dst1)
	// 创建dst2，只复制src的前3个元素

	s1 := []int{1, 2, 3}
	s2 := make([]int, 2)
	copy(s2, s1)
	s2[0] = 999
	fmt.Println(s1)
	fmt.Println(s2)

	sty := []int{1, 2, 3, 4, 5, 6}
	//删除第二个元素
	sty1 := append(sty[:2], sty[3:]...)
	fmt.Println(sty1)

}
