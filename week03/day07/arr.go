package main

import "fmt"

func arr() {
	var arr1 [3]int
	arr1[0] = 21
	arr1[1] = 212
	arr1[2] = 213
	fmt.Println(arr1)

	var numArray = [...]int{1, 2, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(numArray)
}
