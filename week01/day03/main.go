/*
题目一： 姓名（string类型）
年龄（int类型）
身高（float64类型）
是否学生（bool类型）
*/
package main

import "fmt"

/* func getusername() (string, int) {
	return "张三", 19
} */

func getinfo() (string, int, bool) {
	return "张三", 19, true
}

func getStudentInfo() (string, int, string, bool) {
	return "李华", 19, "三年级", true
}
func main() {
	/* 	var name string
	   	var age int
	   	var score = 95.5
	   	var isPass bool = true

	   	name = "李四"
	   	age = 18

	   	fmt.Println(name, age, score, isPass) */

	/*
		 	var username, _ = getusername()
			fmt.Println(username)
	*/

	/* 	name, age, _ := getinfo()
	   	_, _, isPass := getinfo()
	   	fmt.Println(name, age)
	   	fmt.Print(isPass) */

	name, age, _ := getinfo()
	fmt.Println(name, age)
	var _ int = 10
	var _ string = "hello"

	a, _ := 1, 2
	b, _ := 3, 4

	fmt.Println(a, b)

	name, age, _, isPass := getStudentInfo()
	fmt.Println(name, age, isPass)

	const Pi = 3.1415926525

	const VERSION = "v1.0.1"
	fmt.Println(Pi, VERSION)
}
